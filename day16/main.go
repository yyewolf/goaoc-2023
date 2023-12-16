package main

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func st(pos, dir int) uint64 {
	return uint64(pos)<<2 | uint64(dir)
}

var grid = make([]byte, 0, 1000)
var passedThrough = make(map[int]bool)
var memory = make(map[uint64]bool)

func dfs(pos, startDir, width int) int {
	stack := []uint64{st(pos, startDir)}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		dir := int(current & 3)
		pos := int(current >> 2)

		if pos < 0 || pos >= len(grid) {
			continue
		}

		if memory[current] {
			continue
		}
		memory[current] = true
		passedThrough[pos] = true

		x := pos % width

		switch grid[pos] {
		case '.':
			switch dir {
			case 0:
				stack = append(stack, st(pos-width, dir))
			case 1:
				if x == width-1 {
					continue
				}
				stack = append(stack, st(pos+1, dir))
			case 2:
				stack = append(stack, st(pos+width, dir))
			case 3:
				if x == 0 {
					continue
				}
				stack = append(stack, st(pos-1, dir))
			}

		case '|':
			if dir == 1 || dir == 3 {
				stack = append(stack, st(pos-width, 0), st(pos+width, 2))
			} else {
				if dir == 0 {
					stack = append(stack, st(pos-width, dir))
				} else {
					stack = append(stack, st(pos+width, dir))
				}
			}

		case '-':
			if dir == 0 || dir == 2 {
				if x == width-1 {
					stack = append(stack, st(pos-1, 3))
				} else if x == 0 {
					stack = append(stack, st(pos+1, 1))
				} else {
					stack = append(stack, st(pos-1, 3), st(pos+1, 1))
				}
			} else {
				if dir == 1 {
					if x == width-1 {
						continue
					}
					stack = append(stack, st(pos+1, dir))
				} else {
					if x == 0 {
						continue
					}
					stack = append(stack, st(pos-1, dir))
				}
			}

		case '/':
			switch dir {
			case 0:
				if x == width-1 {
					continue
				}
				stack = append(stack, st(pos+1, 1))
			case 1:
				stack = append(stack, st(pos-width, 0))
			case 2:
				if x == 0 {
					continue
				}
				stack = append(stack, st(pos-1, 3))
			case 3:
				stack = append(stack, st(pos+width, 2))
			}

		case '\\':
			switch dir {
			case 0:
				if x == 0 {
					continue
				}
				stack = append(stack, st(pos-1, 3))
			case 1:
				stack = append(stack, st(pos+width, 2))
			case 2:
				if x == width-1 {
					continue
				}
				stack = append(stack, st(pos+1, 1))
			case 3:
				stack = append(stack, st(pos-width, 0))
			}
		}
	}

	return len(passedThrough)
}

func doPartOne(input []byte) int {
	var width int

	grid = grid[:0]
	clear(passedThrough)
	clear(memory)

	for i, c := range input {
		if c == '\n' {
			if width == 0 {
				width = i
			}
		} else {
			grid = append(grid, c)
		}
	}

	return dfs(0, 1, width)
}

func doPartTwo(input []byte) int {
	var width int

	grid = grid[:0]

	for i, c := range input {
		if c == '\n' {
			if width == 0 {
				width = i
			}
		} else {
			grid = append(grid, c)
		}
	}

	var max int

	// Top going down
	for x := 0; x < width; x++ {
		clear(passedThrough)
		clear(memory)
		dfs(x, 2, width)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	// Bottom going up
	for x := 0; x < width; x++ {
		clear(passedThrough)
		clear(memory)
		p := len(grid) - width + x
		dfs(p, 0, width)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
	}

	// Left going right
	var p int
	for p < len(grid) {
		clear(passedThrough)
		clear(memory)
		dfs(p, 1, width)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
		p += width
	}

	// Right going left
	p = width - 1
	for p < len(grid) {
		clear(passedThrough)
		clear(memory)
		dfs(p, 3, width)
		if len(passedThrough) > max {
			max = len(passedThrough)
		}
		p += width
	}

	return max
}
