package main

import (
	"bufio"
	"bytes"
	"runtime"
	"slices"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func doPartOne(input []byte) int {
	buf := bufio.NewReader(bytes.NewReader(input))

	sum := 0
	var seq = make([]int, 0, 100)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		seq = seq[:0]
		var temp int
		var tempSign = 1
		for _, c := range line {
			switch {
			case c == ' ':
				seq = append(seq, tempSign*temp)
				temp = 0
				tempSign = 1
			case c >= '0' && c <= '9':
				temp = temp*10 + int(c-'0')
			case c == '-':
				tempSign = -1
			}
		}
		seq = append(seq, tempSign*temp)

		allZeroes := false
		last := 0

		for !allZeroes && len(seq) > 0 {
			allZeroes = true
			last += seq[len(seq)-1]
			for i := 1; i < len(seq); i++ {
				seq[i-1] = seq[i] - seq[i-1]
				if seq[i-1] != 0 {
					allZeroes = false
				}
			}
			seq = seq[:len(seq)-1]
		}

		sum += last
	}

	return sum
}

func doPartTwo(input []byte) int {
	buf := bufio.NewReader(bytes.NewReader(input))

	var seq = make([]int, 0, 100)

	sum := 0
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		seq = seq[:0]
		var temp int
		var tempSign = 1
		for _, c := range line {
			switch {
			case c == ' ':
				seq = append(seq, tempSign*temp)
				temp = 0
				tempSign = 1
			case c >= '0' && c <= '9':
				temp = temp*10 + int(c-'0')
			case c == '-':
				tempSign = -1
			}
		}
		seq = append(seq, tempSign*temp)

		slices.Reverse(seq)

		stop := false
		last := 0

		for !stop && len(seq) > 0 {
			stop = true
			last += seq[len(seq)-1]
			for i := 1; i < len(seq); i++ {
				seq[i-1] = seq[i] - seq[i-1]
				if seq[i-1] != 0 {
					stop = false
				}
			}
			seq = seq[:len(seq)-1]
		}

		sum += last
	}

	return sum
}
