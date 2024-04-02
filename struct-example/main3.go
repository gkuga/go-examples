package main

import (
	"fmt"
)

func main() {
	var f foo

	f.Add()
	f.Add()
	f.Add()
	go doSomethingPtr1(&f)
	go doSomethingPtr2(&f)

	f.Watch()
	fmt.Println("ダーッ！")
}
