package main

import (
	"bytes"
)

func doPartOne(input []byte) int {
	var startX int = 1
	var startY int = 0

	grid := bytes.Split(input, []byte("\n"))

	var endX int = len(grid[0]) - 2
	var endY int = len(grid) - 1

	// We need to implement a backtracking algorithm
	// Rules :
	// paths (.), forest (#), and steep slopes (^, >, v, and <).
	// You can only go on path or arrows that are not pointing at you
	// You can only go on a path once
	// You cannot go through a forest

	// backtracking algorithm

	var visited [][]bool = make([][]bool, endY+1)
	for i := 0; i < endY+1; i++ {
		visited[i] = make([]bool, endX+1)
	}

	var dfs func(x int, y int, fromX int, fromY int, visited [][]bool) int
	dfs = func(x int, y int, fromX int, fromY int, visited [][]bool) int {
		// check if we are at the end
		if x == endX && y == endY {
			return 1
		}

		// check if we are out of bounds
		if x < 0 || x > endX || y < 0 || y > endY {
			return 0
		}

		// check if we have already been here
		if visited[y][x] {
			return 0
		}

		// check if we are on a forest
		if grid[y][x] == '#' {
			return 0
		}

		// check if we are on a arrow
		if grid[y][x] == '^' || grid[y][x] == '>' || grid[y][x] == 'v' || grid[y][x] == '<' {
			// check if we are going the wrong way
			if (grid[y][x] == '^' && fromY == y-1) || (grid[y][x] == '>' && fromX == x+1) || (grid[y][x] == 'v' && fromY == y+1) || (grid[y][x] == '<' && fromX == x-1) {
				return 0
			}
		}

		// mark as visited
		visited[y][x] = true

		// check if we can go up
		var up int = dfs(x, y-1, x, y, visited)
		// check if we can go right
		var right int = dfs(x+1, y, x, y, visited)
		// check if we can go down
		var down int = dfs(x, y+1, x, y, visited)
		// check if we can go left
		var left int = dfs(x-1, y, x, y, visited)

		// unmark as visited
		visited[y][x] = false

		return max(up, right, down, left) + 1
	}

	return dfs(startX, startY, startX, startY, visited) - 1
}
