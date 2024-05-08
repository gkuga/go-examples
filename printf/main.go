package main

import "fmt"

func main() {
	fmt.Printf("%x\n", []byte{1, 2, 3, 255})
	fmt.Printf("%x\n", 1)
}
