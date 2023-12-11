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

var lineWithGalaxy [150]bool
var columnWithGalaxy [150]bool

var galaxyList = make([][]int, 0, 1000)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func doPartOne(input []byte) int {
	galaxyList = galaxyList[:0]
	var lineOffset = 0
	var y int
	var x int
	for _, c := range input {
		switch c {
		case '\n':
			if !lineWithGalaxy[y] {
				lineOffset += 1
			}
			y++
			x = -1
		case '#':
			galaxyList = append(galaxyList, []int{y + lineOffset, x})
			lineWithGalaxy[y] = true
			columnWithGalaxy[x] = true
		}
		x++
	}

	for index, g := range galaxyList {
		// Adjust coordinates to take into accounts column offsets
		var columnOffset = 0
		for j := 0; j < g[1]; j++ {
			if !columnWithGalaxy[j] {
				columnOffset += 1
			}
		}
		galaxyList[index][1] += columnOffset
	}

	sum := 0
	for i := 0; i < len(galaxyList); i++ {
		for j := i + 1; j < len(galaxyList); j++ {
			// Calculate distance between i and j
			var distance = abs(galaxyList[i][0]-galaxyList[j][0]) + abs(galaxyList[i][1]-galaxyList[j][1])
			sum += distance
		}
	}

	// Safety check for when I'm optimizing :
	if sum != 9543156 {
		fmt.Println("Expected 9543156, got", sum)
	}

	return sum
}

func doPartTwo(input []byte) int {
	galaxyList = galaxyList[:0]
	var lineOffset = 0
	var y int
	var x int
	for _, c := range input {
		switch c {
		case '\n':
			if !lineWithGalaxy[y] {
				lineOffset += 100000 - 1
			}
			y++
			x = -1
		case '#':
			galaxyList = append(galaxyList, []int{y + lineOffset, x})
			lineWithGalaxy[y] = true
			columnWithGalaxy[x] = true
		}
		x++
	}

	for index, g := range galaxyList {
		// Adjust coordinates to take into accounts column offsets
		var columnOffset = 0
		for j := 0; j < g[1]; j++ {
			if !columnWithGalaxy[j] {
				columnOffset += 100000 - 1
			}
		}
		galaxyList[index][1] += columnOffset
	}

	sum := 0
	for i := 0; i < len(galaxyList); i++ {
		for j := i + 1; j < len(galaxyList); j++ {
			// Calculate distance between i and j
			var distance = abs(galaxyList[i][0]-galaxyList[j][0]) + abs(galaxyList[i][1]-galaxyList[j][1])
			sum += distance
		}
	}

	// Safety check for when I'm optimizing :
	// if sum != 625243292686 {
	// 	fmt.Println("Expected 625243292686, got", sum)
	// }

	return sum
}
