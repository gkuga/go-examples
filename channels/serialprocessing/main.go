package main

import (
	"fmt"
)

func main() {
	worker := NewWorker()
	worker.Start()

	// Simulate multiple callers
	for i := 1; i <= 3; i++ {
		resp, err := worker.Send(i)
		if err != nil {
			fmt.Printf("Request %d, Error: %v\n", i, err)
			continue
		}
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
	started chan struct{}
}

// NewWorker function initializes a Worker
func NewWorker() *Worker {
	return &Worker{
		ReqChan: make(chan Request),
		started: make(chan struct{}),
	}
}

// Start method to start the worker
func (w *Worker) Start() {
	go w.run()
	<-w.started
}

func (w *Worker) run() {
	close(w.started)
	for req := range w.ReqChan {
		// Process the request and send response
		resp := req.Value * 2 // Example: double the value
		req.RespChan <- resp  // Send response to the specific channel
	}
}

func (w *Worker) checkStarted() error {
	select {
	case <-w.started:
		return nil
	default:
		return fmt.Errorf("Worker.Start() has not been called")
	}
}

func (w *Worker) Send(value int) (int, error) {
	if err := w.checkStarted(); err != nil {
		return 0, err
	}
	respChan := make(chan int)
	req := Request{Value: value, RespChan: respChan}
	w.ReqChan <- req
	resp := <-respChan
	return resp, nil
}
