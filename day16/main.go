package main

import (
	"bytes"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

var grid [][]byte
var passedThrough = make(map[int]bool)
var memory = make(map[State]bool)

type State struct {
	x, y, dir int
}

func backtrack(startX, startY, startDir int) {
	stack := []State{{startX, startY, startDir}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y, dir := current.x, current.y, current.dir

		if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
			continue
		}

		if memory[current] {
			continue
		}
		memory[current] = true

		passedThrough[y*len(grid[0])+x] = true

		switch grid[y][x] {
		case '.':
			switch dir {
			case 0:
				stack = append(stack, State{x, y - 1, dir})
			case 1:
				stack = append(stack, State{x + 1, y, dir})
			case 2:
				stack = append(stack, State{x, y + 1, dir})
			case 3:
				stack = append(stack, State{x - 1, y, dir})
			}

		case '|':
			if dir == 1 || dir == 3 {
				stack = append(stack, State{x, y - 1, 0}, State{x, y + 1, 2})
			} else {
				if dir == 0 {
					stack = append(stack, State{x, y - 1, dir})
				} else {
					stack = append(stack, State{x, y + 1, dir})
				}
			}

		case '-':
			if dir == 0 || dir == 2 {
				stack = append(stack, State{x - 1, y, 3}, State{x + 1, y, 1})
			} else {
				if dir == 1 {
					stack = append(stack, State{x + 1, y, dir})
				} else {
					stack = append(stack, State{x - 1, y, dir})
				}
			}

		case '/':
			switch dir {
			case 0:
				stack = append(stack, State{x + 1, y, 1})
			case 1:
				stack = append(stack, State{x, y - 1, 0})
			case 2:
				stack = append(stack, State{x - 1, y, 3})
			case 3:
				stack = append(stack, State{x, y + 1, 2})
			}

		case '\\':
			switch dir {
			case 0:
				stack = append(stack, State{x - 1, y, 3})
			case 1:
				stack = append(stack, State{x, y + 1, 2})
			case 2:
				stack = append(stack, State{x + 1, y, 1})
			case 3:
				stack = append(stack, State{x, y - 1, 0})
			}
		}
	}
}

func doPartOne(input []byte) int {
	grid = bytes.Split(input, []byte("\n"))

	backtrack(0, 0, 1)

	return len(passedThrough)
}

func doPartTwo(input []byte) int {

	grid = bytes.Split(input, []byte("\n"))

	var max int

	// Top going down
	for x := range grid[0] {
		clear(passedThrough)
		clear(memory)
		backtrack(x, 0, 2)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	// Bottom going up
	for x := range grid[0] {
		clear(passedThrough)
		clear(memory)
		backtrack(x, len(grid)-1, 0)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	// Left going right
	for y := range grid {
		clear(passedThrough)
		clear(memory)
		backtrack(0, y, 1)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	// Right going left
	for y := range grid {
		clear(passedThrough)
		clear(memory)
		backtrack(len(grid[0])-1, y, 3)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	return max
}
