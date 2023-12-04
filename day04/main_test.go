package main

import "testing"

var result int

func BenchmarkParseCard(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = doCard([]byte("Card 130: 97 57 99 82 46 73 48 25 47 12 | 30  5 77 75 35 67 18 37 52 64 74 38 11 59 41 68 80 73 83 46 71 36 33 84 47"))
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
		r = doPartTwo(input)
	}
	result = r
}
