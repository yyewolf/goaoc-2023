package main

import (
	"testing"
)

var result int

var emptyCards = [1000]Card{}

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copy(cards, emptyCards[:])
		b.StartTimer()

		r = doPartOne(input)
	}
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwo(input)
	}
	result = r
}
