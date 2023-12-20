package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doPartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doPartTwo(input)
	}
}
