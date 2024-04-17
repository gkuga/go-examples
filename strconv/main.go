package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatInt(int64(0), 2))
	fmt.Println(len(strconv.FormatInt(int64(0), 2)))
}
