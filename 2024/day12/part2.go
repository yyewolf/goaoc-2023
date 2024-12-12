package main

import (
	"bytes"
)

var cornerFinder = map[pos][]pos{
	{1, 1}:   []pos{{1, 0}, {0, 1}},
	{-1, 1}:  []pos{{-1, 0}, {0, 1}},
	{1, -1}:  []pos{{1, 0}, {0, -1}},
	{-1, -1}: []pos{{-1, 0}, {0, -1}},
}

func isInBound(i pos, height, width int) bool {
	return i.x >= 0 && i.y >= 0 && i.y < height && i.x < width
}

func match(i1, i2 pos, grid [][]byte) bool {
	maxRow, maxCol := len(grid), len(grid[0])

	if !isInBound(i1, maxRow, maxCol) && !isInBound(i2, maxRow, maxCol) {
		return true
	} else if isInBound(i1, maxRow, maxCol) && isInBound(i2, maxRow, maxCol) {
		p1, p2 := grid[i1.y][i1.x], grid[i2.y][i2.x]
		return p1 == p2
	} else {
		return false
	}
}

func doRegionSides(lines [][]byte, visited map[pos]bool, start pos) (area, sides int) {
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

		for _, neighbor := range neighbors {
			neighborPos := pos{at.x + neighbor[0], at.y + neighbor[1]}
			if !isInBound(neighborPos, height, width) || lines[neighborPos.y][neighborPos.x] != val {
				continue
			}
			stack = append(stack, neighborPos)
		}

		for corner, pair := range cornerFinder {
			c := pos{y: at.y + corner.y, x: at.x + corner.x}
			i1 := pos{y: at.y + pair[0].y, x: at.x + pair[0].x}
			i2 := pos{y: at.y + pair[1].y, x: at.x + pair[1].x}

			if !match(i1, at, lines) && !match(i2, at, lines) {
				sides++
			}
			if match(i1, at, lines) && match(i2, at, lines) && !match(at, c, lines) {
				sides++
			}
		}
	}

	return
}

func doPartTwo(input []byte) int {
	var lines = bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var score int
	var visited = make(map[pos]bool)
	for y, line := range lines {
		for x := range line {
			if visited[pos{x, y}] {
				continue
			}

			area, corners := doRegionSides(lines, visited, pos{x, y})
			score += area * corners
		}
	}

	return score
}
