package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type foo struct {
	n int64
	q chan struct{}

	noCopy noCopy
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

func (f *foo) Add() {
	if atomic.AddInt64(&f.n, 1) == 1 {
		f.q = make(chan struct{})
	}
}

func (f *foo) Done() {
	if atomic.AddInt64(&f.n, -1) == 0 {
		f.q <- struct{}{}
	}
}

func (f *foo) Watch() {
	<-f.q
}

func doSomething1(f foo) {
	time.Sleep(2 * time.Second)
	fmt.Println("さーん！")
	time.Sleep(time.Second)
	f.Done()
}

func doSomething2(f foo) {
	fmt.Println("いーち！")
	time.Sleep(time.Second)
	f.Done()
	fmt.Println("にー！")
	time.Sleep(time.Second)
	f.Done()
}

func doSomethingWithArg1(f foo) {
	time.Sleep(2 * time.Second)
	fmt.Println("さーん！")
	time.Sleep(time.Second)
	f.Done()
}

func doSomethingWithArg2(f foo) {
	fmt.Println("いーち！")
	time.Sleep(time.Second)
	f.Done()
	fmt.Println("にー！")
	time.Sleep(time.Second)
	f.Done()
}

func doSomethingPtr1(f *foo) {
	time.Sleep(2 * time.Second)
	fmt.Println("さーん！")
	time.Sleep(time.Second)
	f.Done()
}

func doSomethingPtr2(f *foo) {
	fmt.Println("いーち！")
	time.Sleep(time.Second)
	f.Done()
	fmt.Println("にー！")
	time.Sleep(time.Second)
	f.Done()
}
