package main

// take returns a channel that yields up to num values from the input channel.
func take[T any](done Done, in <-chan T, num int) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func takeSample() {
	done := make(chan interface{})
	defer close(done)

	for v := range take(done, repeat(done, 1, 2, 3), 5) {
		println(v)
	}
}
