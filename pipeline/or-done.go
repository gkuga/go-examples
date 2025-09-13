package main

func orDone[T any](done Done, c <-chan T) <-chan T {
	valStream := make(chan T)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func multiplyOrDone(done Done, intStream Stream, multiplier int) Stream {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for n := range orDone(done, intStream) {
			multipliedStream <- n * multiplier
		}
	}()
	return multipliedStream
}
