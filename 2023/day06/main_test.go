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

func BenchmarkPartTwo(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwo(input)
	}
	result = r
}

func BenchmarkPartOneMath(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartOneMath(input)
	}
	result = r
}

func BenchmarkPartTwoMath(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwoMath(input)
	}
	result = r
}
