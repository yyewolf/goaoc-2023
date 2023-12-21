package main

import (
	"bytes"
	"fmt"
	"math"
)

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func BFS(grid [][]byte, startX, startY int, maxSteps int) int {
	var attainable int
	var visited = make(map[string]bool)
	var parity = maxSteps % 2

	// BFS queue
	var queue [][3]int

	queue = append(queue, [3]int{startX, startY, 0})

	for len(queue) > 0 {
		// pop first element
		current := queue[0]
		queue = queue[1:]

		x := current[0]
		y := current[1]
		steps := current[2]

		if steps > maxSteps {
			continue
		}

		key := fmt.Sprintf("%d,%d", x, y)
		if visited[key] {
			continue
		}

		visited[key] = true

		if steps%2 == parity {
			attainable++
		}

		// add neighbors to queue
		for _, neighbor := range [][2]int{
			{x - 1, y},
			{x + 1, y},
			{x, y - 1},
			{x, y + 1},
		} {
			nx := neighbor[0]
			ny := neighbor[1]

			// if nx < 0 || ny < 0 || nx >= len(grid[0]) || ny >= len(grid) {
			// 	continue
			// }
			// We loop infinitely, if x goes below 0, it wraps around to len(grid[0]) - 1
			if grid[mod(ny, len(grid))][mod(nx, len(grid[0]))] == '#' {
				continue
			}

			k := fmt.Sprintf("%d,%d", nx, ny)
			if visited[k] {
				continue
			}

			queue = append(queue, [3]int{nx, ny, steps + 1})
		}
	}

	return attainable
}

func doPartTwo(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))

	startX := 0
	startY := 0

	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				startX = x
				startY = y
			}
		}
	}

	mod := 26501365 % len(grid)

	// BFS from start to find each position attainable in exactly 64 steps
	a := BFS(grid, startX, startY, mod)
	b := BFS(grid, startX, startY, mod+len(grid))
	c := BFS(grid, startX, startY, mod+len(grid)+len(grid))

	ba := b - a
	cb := c - b
	cbba := cb - ba

	A := cbba / 2
	B := ba - 3*A
	C := a - B - A

	x := int(math.Ceil(26501365.0 / float64(len(grid))))

	attainable := A*x*x + B*x + C

	return attainable
}
