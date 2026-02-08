package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ticker, err := NewFileTicker(ctx, 10*time.Second, "tick_info.json")
	if err != nil {
		fmt.Println("Error creating file ticker:", err)
		cancel()
		return
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		timeout := time.After(20 * time.Second)
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case t := <-ticker.C:
				if t.err != nil {
					fmt.Println("Error during tick:", t.err)
					continue
				}
				fmt.Println("Tick at", t.time)
			case <-signalChan:
				cancel()
				fmt.Println("Received interrupt signal, stopping ticker...")
			case <-timeout:
				cancel()
				fmt.Println("Timeout reached, stopping ticker...")
			}
		}
	}()
	wg.Wait()
	fmt.Println("Ticker stopped")
}
