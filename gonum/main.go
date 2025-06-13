package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func example() {
	A := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})
	v := mat.NewVecDense(2, []float64{5, 6})

	var w mat.VecDense
	w.MulVec(A, v)

	fmt.Printf("w = %v\n", w.RawVector().Data)
}

func main() {
	example()
}
