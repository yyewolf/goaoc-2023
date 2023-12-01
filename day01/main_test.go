package main

import "testing"

var result int

func BenchmarkPartOne(b *testing.B) {
	file, err := parse("input.txt")
	if err != nil {
		panic(err)
	}
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartOne(file)
	}
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	file, err := parse("input.txt")
	if err != nil {
		panic(err)
	}
	var r int
	for n := 0; n < b.N; n++ {
		r = doPartTwo(file)
	}
	result = r
}

func BenchmarkPartOneParsing(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		file, err := parse("input.txt")
		if err != nil {
			panic(err)
		}
		r = doPartOne(file)
	}
	result = r
}

func BenchmarkPartTwoParsing(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		file, err := parse("input.txt")
		if err != nil {
			panic(err)
		}
		r = doPartTwo(file)
	}
	result = r
}
