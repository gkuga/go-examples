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

func tickChannel() {
	ch := make(chan time.Time)
	go func() {
		for t := range time.Tick(time.Second) {
			ch <- t
		}
	}()
	go func() {
		for t := range ch {
			fmt.Println("受信:", t)
		}
	}()
	select {}
}

func afterChannel() {
	ch := make(chan bool)
	go func() {
		for {
			ch <- true
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			<-ch
			fmt.Println("受信:", time.Now())
		}
	}()
	select {}
}

func emptySelectDeadLock1() {
	select {}
}

// select{}でgoroutineはasleepステータスに移行する？
func emptySelectDeadLock2() {
	go func() {
		select {}
	}()
	select {}
}

func emptySelectNotDeadLock() {
	go func() {
		for {
		}
	}()
	select {}
}

func main() {
	emptySelectDeadLock2()
}
