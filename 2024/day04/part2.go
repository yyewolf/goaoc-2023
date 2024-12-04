package main

import (
	"bytes"
)

// check if the A is in a X shape MAS like this :
// M . S
// . A .
// M . S
//
// M must be adjacent to another M that implies the same for S

var masPositions = [][]int{
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func isXMas(i, j int, lines [][]byte) bool {
	var square [3][3]byte
	var mapCount = make(map[byte]int)
	for _, d := range masPositions {
		x := i + d[0]
		y := j + d[1]

		if x < 0 || x >= len(lines) || y < 0 || y >= len(lines[x]) {
			continue
		}

		mapCount[lines[x][y]] += 1
		square[d[0]+1][d[1]+1] = lines[x][y]
	}

	if mapCount['M'] != 2 || mapCount['S'] != 2 {
		return false
	}

	// Verify that upper left and lower right are not the same
	if square[0][0] == square[2][2] {
		return false
	}

	return true
}

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var count int
	for i, line := range lines {
		for j, char := range line {
			if char == 'A' && isXMas(i, j, lines) {
				count += 1
			}
		}
	}

	return count
}
