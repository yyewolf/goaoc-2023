package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		clear(passedThrough)
		clear(memory)
		b.StartTimer()
		doPartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doPartTwo(input)
	}
}
