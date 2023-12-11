package main

import (
	"fmt"
	"runtime"
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var galaxyList [500]uint64
var columnWithGalaxy [140]int

func doIt(offset int, input []byte) int {
	var lineWithGalaxy bool
	var lineOffset = 0
	var y int
	var x int
	var galaxyCount int
	for _, c := range input {
		switch c {
		case '\n':
			if !lineWithGalaxy {
				lineOffset += 1
			}
			y++
			x = -1
			lineWithGalaxy = false
		case '#':
			// Put line offset in the uppermost 2 bits
			galaxyList[galaxyCount] = uint64(y+lineOffset*(offset-1))<<32 | uint64(x)
			lineWithGalaxy = true
			columnWithGalaxy[x] = 1
			galaxyCount++
		}
		x++
	}

	var colWithGalaxy [140]int

	cum := 0
	for i, c := range columnWithGalaxy {
		if c == 0 {
			cum++
		}
		colWithGalaxy[i] = cum
	}

	for i := 0; i < galaxyCount; i++ {
		g := galaxyList[i]
		// Adjust coordinates to take into accounts column offsets
		x := int(g & 0xFFFFFF)
		columnOffset := colWithGalaxy[x]
		galaxyList[i] += uint64(columnOffset * (offset - 1))
	}

	// [(0, 4), (1, 9), (2, 0), (5, 8), (6, 1), (7, 12), (10, 9), (11, 0), (11, 5)]

	sum := 0
	for i := 0; i < galaxyCount; i++ {
		x, y := int(galaxyList[i]&0xFFFFFF), int(galaxyList[i]>>32)
		for j := i + 1; j < galaxyCount; j++ {
			x1, y1 := int(galaxyList[j]&0xFFFFFF), int(galaxyList[j]>>32)

			// Calculate distance between i and j
			var distance = abs(x-x1) + abs(y-y1)
			sum += distance
		}
	}

	return sum
}

func doPartOne(input []byte) int {
	sum := doIt(2, input)

	// Safety check for when I'm optimizing :
	if sum != 9543156 {
		fmt.Println("Expected 9543156, got", sum)
	}

	return sum
}

func doPartTwo(input []byte) int {
	sum := doIt(1000000, input)

	// Safety check for when I'm optimizing :
	if sum != 625243292686 {
		fmt.Println("Expected 625243292686, got", sum)
	}

	return sum
}
