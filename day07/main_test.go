package main

import (
	"testing"
)

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		clear(cards)
		b.StartTimer()

		r = doPartOne(input)
	}
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		clear(cards)
		b.StartTimer()

		r = doPartTwo(input)
	}
	result = r
}
