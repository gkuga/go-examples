package main

import (
	"fmt"
)

func main() {
	var f foo

	f.Add()
	f.Add()
	f.Add()
	go doSomethingWithArg1(f)
	go doSomethingWithArg2(f)

	f.Watch()
	fmt.Println("ダーッ！")
}
