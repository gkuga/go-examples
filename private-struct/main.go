package main

import (
	"fmt"

	"github.com/gkuga/go-examples/private-struct/foo"
)

func main() {
	//var a *foo.bar = foo.Newbar(123) // doesn't work, because the name 'bar' is not exported
	a := foo.Newbar(123)
	a.Val = 100

	fmt.Println(a.Val) // prints 123
}
