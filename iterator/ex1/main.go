package main

import (
	"fmt"
	"iter"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for chunk := range Chunk(nums, 3) {
		fmt.Println(chunk)
	}
}

func ChunkWithGenerics[S ~[]T, T any](slice S, size int) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}

	var chunks [][]T
	for i := 0; i <= len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice] {
	if n < 1 {
		panic("cannot be less than 1")
	}

	return func(yield func(Slice) bool) {
		for i := 0; i < len(s); i += n {
			end := min(n, len(s[i:]))
			if !yield(s[i : i+end : i+end]) {
				return
			}
		}
	}
}
