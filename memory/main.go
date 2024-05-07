package main

import (
	"fmt"
	"runtime"
	"time"
)

var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}

type MyStruct struct {
	Data []byte
}

func process1() {
	var data []*MyStruct
	for i := 0; i < 10000; i++ {
		data = append(data, &MyStruct{Data: make([]byte, 1000)})
	}
}

func process2() {
	var data []*MyStruct
	for i := 0; i < 10000; i++ {
		data = append(data, &MyStruct{Data: make([]byte, 1000)})
	}
	for i := range data {
		data[i] = nil
	}
}

func main() {
	PrintMemory()
	process1()
	for {
		time.Sleep(time.Second * 5)
		PrintMemory()
	}
}
