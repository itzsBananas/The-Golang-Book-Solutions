package main

import (
	"testing"

	"gopl.io/ch2/popcount"
)

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkShiftCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shiftCount(0x1234567890ABCDEF)
	}
}
