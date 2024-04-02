package main

import (
	"fmt"
	"time"
)

func main() {
	var f foo

	f.Add()
	f.Add()
	f.Add()
	go func() {
		fmt.Println("いーち！")
		time.Sleep(time.Second)
		f.Done()
		fmt.Println("にー！")
		time.Sleep(time.Second)
		f.Done()
		fmt.Println("さーん！")
		time.Sleep(time.Second)
		f.Done()
	}()

	f.Watch()
	fmt.Println("ダーッ！")
}
