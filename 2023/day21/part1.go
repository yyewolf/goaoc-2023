package main

import (
	"bytes"
)

func doPartOne(input []byte) int {
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

	// BFS from start to find each position attainable in exactly 64 steps

	var attainable [150][150]bool
	var visited [150][150]bool

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

		if visited[y][x] {
			continue
		}

		visited[y][x] = true

		if steps == 64 {
			attainable[y][x] = true
			continue
		}

		if steps%2 == 0 {
			attainable[y][x] = true
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

			if nx < 0 || ny < 0 || nx >= len(grid[0]) || ny >= len(grid) {
				continue
			}

			if grid[ny][nx] == '#' {
				continue
			}

			if visited[ny][nx] {
				continue
			}

			queue = append(queue, [3]int{nx, ny, steps + 1})
		}
	}

	// // print grid with O if attainable, X if not
	// for y, row := range grid {
	// 	for x, cell := range row {
	// 		if cell == '#' {
	// 			fmt.Print("#")
	// 			continue
	// 		}

	// 		if attainable[y][x] {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	s := 0
	for _, row := range attainable {
		for _, cell := range row {
			if cell {
				s++
			}
		}
	}

	return s
}
