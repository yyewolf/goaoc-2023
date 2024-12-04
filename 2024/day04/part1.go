package main

import (
	"bytes"
)

const (
	// Word is MAS to avoid counting X twice
	word = "MAS"
)

var directions = [][]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

// checkWord checks following the given direction
func checkWord(i, j int, d []int, lines [][]byte) bool {
	for k := 0; k < len(word); k++ {
		i += d[0]
		j += d[1]

		if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
			return false
		}

		if lines[i][j] != word[k] {
			return false
		}
	}

	return true
}

// countWords count the words when reaching an X, it proceeds to count in every directions
func countWords(i, j int, lines [][]byte) int {
	var count int
	for _, d := range directions {
		if checkWord(i, j, d, lines) {
			count++
		}
	}

	return count
}

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var count int
	for i, line := range lines {
		for j, char := range line {
			if char == 'X' {
				c := countWords(i, j, lines)
				count += c
			}
		}
	}

	return count
}
