package main

import (
	"fmt"
)

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) Ok() bool {
	return r.err == nil
}

func (r Result[T]) Value() T {
	if r.Ok() {
		return r.value
	}
	var def T
	return def
}

func main() {
	res := Result[string]{value: "Hello World!"}
	fmt.Printf("res: %v\n", res.Value())
	res.err = fmt.Errorf("Error!")
	fmt.Printf("res: %v\n", res.Value())
	res.err = nil
	fmt.Printf("res: %v\n", res.Value())
}
