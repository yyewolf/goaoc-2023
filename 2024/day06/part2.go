package main

import (
	"bytes"
	"slices"
)

func doPartTwo(input []byte) int {
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
	visited := make(map[[2]int][]int)

	// the guard goes forward until it hits a #, then turns right by 90 degrees
	var direction int

	for {
		visited[[2]int{guardX, guardY}] = append(visited[[2]int{guardX, guardY}], direction)

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

	// Find all the positions where we can insert a new obstruction to make the guard go in a loop
	//

	// Function to simulate guard's path and check if it loops
	isLooping := func(grid [][]byte, guardX, guardY int, obstruction [2]int) bool {
		direction := 0 // start facing up
		visited := make(map[[2]int][]int)

		for {
			// Check if the guard is stuck in a loop
			if slices.Contains(visited[[2]int{guardX, guardY}], direction) {
				return true
			}
			visited[[2]int{guardX, guardY}] = append(visited[[2]int{guardX, guardY}], direction)

			// Try to move forward
			nextX := guardX + directions[direction][0]
			nextY := guardY + directions[direction][1]

			if nextX < 0 || nextX >= len(grid[0]) || nextY < 0 || nextY >= len(grid) {
				return false
			}

			// Check for out of bounds or hitting a # or the obstruction
			if grid[nextY][nextX] == '#' || (nextX == obstruction[0] && nextY == obstruction[1]) {
				// Turn right
				direction = (direction + 1) % 4
			} else {
				// Move forward
				guardX, guardY = nextX, nextY
			}
		}
	}

	// find the guard's position
	guardX, guardY = 0, 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == guard {
				guardX = x
				guardY = y
				break
			}
		}
	}

	// Try placing an obstruction in every empty space
	validPositions := 0
	for pos, _ := range visited {
		x, y := pos[0], pos[1]
		if grid[y][x] == '.' && (x != guardX || y != guardY) {
			// Test placing an obstruction at (x, y)
			if isLooping(grid, guardX, guardY, [2]int{x, y}) {
				validPositions++
			}
		}
	}

	// count the visited cells
	return validPositions
}
