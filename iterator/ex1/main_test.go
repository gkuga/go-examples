package main

import (
	"slices"
	"testing"
)

func BenchmarkChunk_BySlicePkg(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	size := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slices.Chunk(nums, size)
	}
}

func BenchmarkChunk_ByGenerics(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	size := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ChunkWithGenerics(nums, size)
	}
}
