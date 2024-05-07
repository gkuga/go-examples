package main

import (
	"fmt"
	"sync"
	"time"
)

func closeChan(ch chan int, sec *int) {
	var d = time.Second * 3
	if sec != nil {
		d = time.Second * time.Duration(*sec)
	}
	time.Sleep(d)
	close(ch)
}

func mainDeadLock1() {
	ch := make(chan int, 2)
  ch <- 1
  ch <- 2
	for v := range ch {
		fmt.Println(v)
	}
}

func mainDeadLock2() {
	ch := make(chan int, 2)
  ch <- 1
  ch <- 2
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}

func mainNotDeadLock1() {
	ch := make(chan int, 2)
  ch <- 1
  ch <- 2
	go closeChan(ch, nil)
	for v := range ch {
		fmt.Println(v)
	}
}

func bufferdChan() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println("cap:", cap(ch), "len:", len(ch))
	<-ch
	fmt.Println("cap:", cap(ch), "len:", len(ch))
}

func selectChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	sec := 3
	go closeChan(ch, &sec)
	for {
		select {
		case v, ok := <-ch:
			fmt.Println(v)
			if !ok {
				return
			}
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("default")
		}
	}
}

func main() {
	selectChan()
	// mainDeadLock1()
	// mainDeadLock2()
}
