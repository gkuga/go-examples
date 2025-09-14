package main

func tee[T any](done Done, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range orDone(done, in) {
			o1, o2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				/*
				   In Go, if a channel in a select case is nil, that case is always blocked and will never be selected.
				   By setting out1 or out2 to nil after sending, we dynamically remove that case from the select candidates.
				   This ensures that each value is sent to both out1 and out2 exactly once.
				*/
				case o1 <- v:
					o1 = nil
				case o2 <- v:
					o2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func teeSample() {
	done := make(chan interface{})
	defer close(done)

	nums := []int{1, 2, 3, 4, 5}
	intStream := generator(done, nums...)
	out1, out2 := tee(done, intStream)

	for {
		v1, ok1 := <-out1
		v2, ok2 := <-out2
		if !ok1 && !ok2 {
			break
		}
		if ok1 {
			println("out1:", v1)
		}
		if ok2 {
			println("out2:", v2)
		}
	}
}
