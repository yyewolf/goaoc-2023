package main

import (
	_ "embed"
	"os"
)

//go:embed input.txt
var input []byte

//go:embed input_test.txt
var inputTest []byte

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}
