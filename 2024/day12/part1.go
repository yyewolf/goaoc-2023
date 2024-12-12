package main

import (
	"bytes"
)

type pos struct {
	x, y int
}

var neighbors = [][2]int{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

func doRegion(lines [][]byte, visited map[pos]bool, start pos) (area, perimeter int) {
	var val = lines[start.y][start.x]
	var stack = []pos{start}

	var height = len(lines)   // y axis
	var width = len(lines[0]) // x axis

	for len(stack) > 0 {
		at := stack[0]
		stack = stack[1:]

		if visited[at] {
			continue
		}
		visited[at] = true
		area += 1
		perimeter += 4

		for _, neighbor := range neighbors {
			neighborPos := pos{at.x + neighbor[0], at.y + neighbor[1]}
			if neighborPos.x < 0 || neighborPos.x > width-1 || neighborPos.y < 0 || neighborPos.y > height-1 {
				continue
			}
			if lines[neighborPos.y][neighborPos.x] != val {
				continue
			}
			stack = append(stack, neighborPos)
			perimeter -= 1
		}
	}
	return
}

func doPartOne(input []byte) int {
	var lines = bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var score int
	var visited = make(map[pos]bool)
	for y, line := range lines {
		for x := range line {
			if visited[pos{x, y}] {
				continue
			}

			area, perimeter := doRegion(lines, visited, pos{x, y})
			score += area * perimeter
		}
	}

	return score
}
