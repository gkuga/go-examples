package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	f := func() {
		fmt.Println("Done")
	}
	once.Do(f)
	once.Do(f)
	once.Do(f)

	// async calls
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(f)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
