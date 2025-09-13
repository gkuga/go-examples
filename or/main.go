package main

import (
	"fmt"
	"time"
)

// or function takes multiple channels and returns a single channel that closes
// when any of the input channels close.

type Done <-chan struct{}

func or(channels ...Done) Done {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan struct{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], Done(orDone))...):
			}
		}
	}()
	return Done(orDone)
}

func main() {
	sig := func(after time.Duration) Done {
		c := make(chan struct{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return Done(c)
	}

	start := time.Now()
	// Breakpoint: before waiting for or-channel
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
