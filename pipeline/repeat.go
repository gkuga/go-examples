package main

// repeat returns a channel that emits the given values infinitely (until done is closed).
func repeat[T any](done Done, values ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()
	return out
}

func repeatSample() {
	done := make(chan interface{})
	defer close(done)

	for v := range take(done, repeat(done, "A", "B"), 4) {
		println(v)
	}
}
