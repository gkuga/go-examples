package main

import "fmt"

type Done <-chan interface{}
type Stream <-chan int

func generator(done Done, nums ...int) Stream {
	intStream := make(chan int, len(nums))
	go func() {
		defer close(intStream)
		for _, n := range nums {
			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()
	return intStream
}

func multiply(done Done, intStream Stream, multiplier int) Stream {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for n := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- n * multiplier:
			}
		}
	}()
	return multipliedStream
}

func add(done Done, intStream Stream, additive int) Stream {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for n := range intStream {
			select {
			case <-done:
				return
			case addedStream <- n + additive:
			}
		}
	}()
	return addedStream
}

func multiplySample() {
	// This is just a placeholder main function.
	done := make(chan interface{})
	defer close(done)

	nums := []int{1, 2, 3, 4, 5}
	intStream := generator(done, nums...)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for result := range pipeline {
		fmt.Println(result)
	}
}

func main() {
	multiplySample()
	orDoneSample()
	bridgeSample()
	teeSample()
	takeSample()
	repeatSample()
}
