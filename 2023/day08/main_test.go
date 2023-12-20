package main

import (
	"testing"
)

var result int

func BenchmarkParsing(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = int(doParsing(input))
	}
	result = r
}

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartOne(input)
	}
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwo(input, b)
	}
	result = r
}
