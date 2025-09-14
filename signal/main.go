package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a channel to receive OS signals
	ch := make(chan os.Signal, 1)
	// Notify the channel on SIGINT and SIGTERM
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Program is running... Press Ctrl+C to exit.")

	// Wait for a signal to be received
	sig := <-ch
	fmt.Printf("Received signal: %v\n", sig)
	fmt.Println("Exiting gracefully.")

	// Uncomment the following line to see the behavior with a zero-buffered channel
	bufferZeroSample()
}

func bufferZeroSample() {
	// Unbuffered channel (buffer size 0)
	// https://github.com/golang/go/blob/ac803b5949f6dbc5bfa559afe506d35f9e1b3195/src/os/signal/signal.go#L233
	ch := make(chan os.Signal)

	// Notify the channel on SIGINT and SIGTERM
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Sleeping for 10 seconds. Try sending SIGINT or SIGTERM now.")
	time.Sleep(10 * time.Second)

	fmt.Println("Now waiting for signal...")
	sig := <-ch
	fmt.Printf("Received signal: %v\n", sig)
	fmt.Println("Exiting gracefully.")
}
