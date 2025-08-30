package main

import (
	"fmt"
)

func main() {
	reqChan := make(chan Request)

	// Create and start Worker
	worker := Worker{ReqChan: reqChan}
	go worker.Run()

	// Simulate multiple callers
	for i := 1; i <= 3; i++ {
		respChan := make(chan int)
		req := Request{Value: i, RespChan: respChan}
		reqChan <- req     // Send request with response channel
		resp := <-respChan // Receive response from worker
		fmt.Printf("Request %d, Response %d\n", i, resp)
	}
}

type Request struct {
	Value    int
	RespChan chan int
}

// Worker struct holds the request channel
type Worker struct {
	ReqChan chan Request
}

// Run method processes incoming requests
func (w *Worker) Run() {
	for req := range w.ReqChan {
		// Process the request and send response
		resp := req.Value * 2 // Example: double the value
		req.RespChan <- resp  // Send response to the specific channel
	}
}

func worker(reqChan chan Request) {
	for req := range reqChan {
		// Process the request and send response
		resp := req.Value * 2 // Example: double the value
		req.RespChan <- resp  // Send response to the specific channel
	}
}
