package main

import "testing"

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartOne(input)
	}
	result = r
}

func BenchmarkPartTwoParsing(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwo(input)
	}
	result = r
}
