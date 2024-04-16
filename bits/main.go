package main

import "fmt"

func main() {
	fmt.Printf("%v\n", uint8(1<<0))

	for i := 1; i < 1<<8; i <<= 1 {
		fmt.Printf("%v\n", i)
	}

	// Compile error
	// fmt.Printf("%v\n", uint8(1<<8))
	for i := 0; i <= 8; i++ {
		fmt.Printf("%v\n", uint8(1<<i))
	}
}
