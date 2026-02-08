package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultFileTickerDuration = time.Second * 1
	defaultFilePermission     = 0644
)

type Tick struct {
	time time.Time
	err  error
}

type FileTicker struct {
	ticker *time.Ticker
	info   info
	ftd    time.Duration
	d      time.Duration
	C      <-chan Tick
}

type info struct {
	NextTick time.Time `json:"next_tick"`
}

func WithFileTickerDuration(d time.Duration) Option {
	return func(t *FileTicker) {
		t.ftd = d
	}
}

type Option func(*FileTicker)

func NewFileTicker(
	ctx context.Context,
	d time.Duration,
	file string,
	opts ...Option,
) (*FileTicker, error) {
	ft := &FileTicker{
		d:   d,
		ftd: defaultFileTickerDuration,
		info: info{
			NextTick: time.Now().Add(d),
		},
	}
	for _, opt := range opts {
		opt(ft)
	}
	err := ft.readFile(file)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	go func() {
		c := make(chan Tick)
		ft.C = c
		ft.start(ctx, file, c)
	}()
	if !os.IsNotExist(err) {
		return ft, nil
	}
	err = ft.createFile(file)
	if err != nil {
		return nil, err
	}
	return ft, nil
}

func (ft *FileTicker) Stop() {
	ft.ticker.Stop()
}

func (ft *FileTicker) start(
	ctx context.Context,
	file string,
	c chan<- Tick,
) {
	ft.ticker = time.NewTicker(ft.ftd)
	defer ft.ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ft.ticker.C:
			if t.Before(ft.info.NextTick) {
				continue
			}
			nextTick := t.Add(ft.d)
			c <- Tick{time: t}
			ft.info.NextTick = nextTick
			if err := ft.updateFile(file); err != nil {
				c <- Tick{err: err}
			}
		}
	}
}

func (ft *FileTicker) readFile(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &ft.info)
	return err
}

func (ft *FileTicker) createFile(file string) error {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.MarshalIndent(ft.info, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile(file, data, defaultFilePermission)
	return nil
}

func (ft *FileTicker) updateFile(file string) error {
	data, err := json.MarshalIndent(ft.info, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, defaultFilePermission)
}
