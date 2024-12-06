package main

import (
	"bytes"
)

const guard = '^'

var directions = [4][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func doPartOne(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))
	grid = grid[:len(grid)-1]

	// find the guard's position
	var guardX, guardY int
	for y, row := range grid {
		for x, cell := range row {
			if cell == guard {
				guardX = x
				guardY = y
				break
			}
		}
	}

	// map the visited cells
	visited := make(map[[2]int]bool)

	// the guard goes forward until it hits a #, then turns right by 90 degrees
	var direction int

	for {
		visited[[2]int{guardX, guardY}] = true

		// check if we can go forward
		nextX := guardX + directions[direction][0]
		nextY := guardY + directions[direction][1]

		// if we're out of the grid, break
		if nextX < 0 || nextX >= len(grid[0]) || nextY < 0 || nextY >= len(grid) {
			break
		}

		if grid[nextY][nextX] != '#' {
			guardX = nextX
			guardY = nextY
		} else {
			// turn right
			direction = (direction + 1) % 4
		}
	}

	// count the visited cells
	return len(visited)
}
