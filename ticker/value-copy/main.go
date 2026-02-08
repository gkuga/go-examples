package main

import (
	"fmt"
	"time"
)

type ticker struct {
	ticker time.Ticker
}

func main() {
	var t1 ticker
	// var err error
	fmt.Printf("Initial ticker address: %p\n", &t1.ticker)
	go func() {
		t1.start()
		fmt.Printf("Ticker address: %p\n", &t1)
		// fmt.Printf("%p\n", &err)
	}()
	defer func() {
		fmt.Printf("Ticker address before stopping: %p\n", &t1.ticker)
		fmt.Printf("Ticker address: %p\n", &t1)
		t1.stop()
		fmt.Println("Ticker stopped successfully.")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed.")
}

func (t *ticker) start() {
	fmt.Println("Initializing ticker...")
	t.ticker = *time.NewTicker(1 * time.Second)
	fmt.Printf("Ticker initialized with address: %p\n", &t.ticker)
	go func() {
		for range t.ticker.C {
			fmt.Println("Ticker ticked.")
		}
	}()
}

func (t *ticker) stop() {
	fmt.Printf("Stopping ticker with address: %p\n", &t.ticker)
	t.ticker.Stop()
	fmt.Println("Ticker stop method executed.")
}
