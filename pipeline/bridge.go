package main

// bridge flattens a channel of channels into a single channel.
// It takes a channel that emits channels, and emits all values from each inner channel in order.
func bridge[T any](done Done, chanStream <-chan <-chan T) <-chan T {
	valStream := make(chan T)
	go func() {
		defer close(valStream)
		for innerChan := range orDone(done, chanStream) {
			for v := range orDone(done, innerChan) {
				valStream <- v
			}
		}
	}()
	return valStream
}

// Sample usage of bridge
// This will flatten a channel of channels into a single channel and print all values.
func bridgeSample() {
	done := make(chan interface{})
	defer close(done)

	// Create a channel of channels
	chanStream := make(chan (<-chan int))
	go func() {
		defer close(chanStream)
		for i := 0; i < 3; i++ {
			ch := make(chan int, 2)
			ch <- i * 10
			ch <- i*10 + 1
			close(ch)
			chanStream <- ch
		}
	}()

	for v := range bridge(done, chanStream) {
		println(v)
	}
}
