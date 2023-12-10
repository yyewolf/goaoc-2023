package main

import (
	"bufio"
	"bytes"
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

var grid [][]byte
var visited []bool = make([]bool, 20000)
var distances []int = make([]int, 20000)

func xyToIndex(x, y int) int {
	return y*len(grid[0]) + x
}

func dfs(x, y, d int) {
	// If we have already visited this node, return
	if visited[xyToIndex(x, y)] {
		return
	}

	// Mark this node as visited
	visited[xyToIndex(x, y)] = true
	distances[xyToIndex(x, y)] = d

	// If this node is a -, we can go to the node on the left or right
	if grid[y][x] == '-' {
		// Go to the node on the left
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		// Go to the node on the right
		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}
	}

	// If this node is a |, we can go to the node above or below
	if grid[y][x] == '|' {
		// Go to the node above
		if y > 0 {
			dfs(x, y-1, d+1)
		}

		// Go to the node below
		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}
	}

	// If this node is a L, we can go to the node on the right or above
	if grid[y][x] == 'L' {
		// Go to the node on the right
		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}

		// Go to the node above
		if y > 0 {
			dfs(x, y-1, d+1)
		}
	}

	// If this node is a J, we can go to the node on the left or above
	if grid[y][x] == 'J' {
		// Go to the node on the left
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		// Go to the node above
		if y > 0 {
			dfs(x, y-1, d+1)
		}

	}

	// If this node is a 7, we can go to the node on the left or below
	if grid[y][x] == '7' {
		// Go to the node on the left
		if x > 0 {
			dfs(x-1, y, d+1)
		}

		// Go to the node below
		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}

	}

	// If this node is a F, we can go to the node on the right or below
	if grid[y][x] == 'F' {
		// Go to the node on the right
		if x < len(grid[y])-1 {
			dfs(x+1, y, d+1)
		}

		// Go to the node below
		if y < len(grid)-1 {
			dfs(x, y+1, d+1)
		}
	}
}

func doPartOne(input []byte) int {
	var b = bufio.NewReader(bytes.NewBuffer(input))

	var startx, starty int

	grid = grid[:0]

	var y int
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		if startx == 0 && starty == 0 {
			for x, c := range line {
				if c == 'S' {
					startx = x
					starty = y
					line[x] = '-'
					break
				}
			}
		}

		grid = append(grid, line[:len(line)-1])

		y++
	}

	// Implement BFS to find the farthest node from the start
	// We start from the start node, and we assign all nodes we can go to a distance of 1
	// Then we go to those nodes, and assign all nodes we can go to a distance of 2
	// We can determine neighbors this way :
	//   If the node is a -, we can go to the node on the left or right
	//   If the node is a |, we can go to the node above or below
	//   If the node is a L, we can go to the node on the right or above
	//   If the node is a J, we can go to the node on the left or above
	//   If the node is a 7, we can go to the node on the left or below
	//   If the node is a F, we can go to the node on the right or below

	// Reset the arrays
	// clear(visited)
	// clear(distances)

	dfs(startx, starty, 0)

	// Find the farthest node
	max := 0
	for _, d := range distances {
		if d > max {
			max = d
		}
	}

	return max/2 + 1
}

func doPartTwo(input []byte) int {
	var b = bufio.NewReader(bytes.NewBuffer(input))

	grid = grid[:0]

	var startx, starty int

	var y int
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}

		if startx == 0 && starty == 0 {
			for x, c := range line {
				if c == 'S' {
					startx = x
					starty = y
					line[x] = '-'
					break
				}
			}
		}

		grid = append(grid, line[:len(line)-1])

		y++
	}

	// Implement BFS to find the farthest node from the start
	// We start from the start node, and we assign all nodes we can go to a distance of 1
	// Then we go to those nodes, and assign all nodes we can go to a distance of 2
	// We can determine neighbors this way :
	//   If the node is a -, we can go to the node on the left or right
	//   If the node is a |, we can go to the node above or below
	//   If the node is a L, we can go to the node on the right or above
	//   If the node is a J, we can go to the node on the left or above
	//   If the node is a 7, we can go to the node on the left or below
	//   If the node is a F, we can go to the node on the right or below

	// Reset the visited array
	// clear(visited)
	// clear(distances)

	// Start the DFS
	dfs(startx, starty, 0)

	// We need to count the amount of nodes that are enclosed by the path
	// We can use the visited array for this
	// Example :
	// ...........
	// .S-------7.
	// .|F-----7|.
	// .||.....||.
	// .||.....||.
	// .|L-7.F-J|.
	// .|..|.|..|.
	// .L--J.L--J.
	// ...........

	// There's only four nodes that are enclosed by the path :
	// .|XX|.|XX|.  because they are surrounded by the path on all sides

	count := 0
	var state bool
	var lastChar byte
	for y, line := range grid {
		for x := range line {
			v := visited[xyToIndex(x, y)]
			if v {
				// Flip the state if we crossed a line
				switch grid[y][x] {
				case '|':
					state = !state
				case '-':
				default:
					if lastChar == '.' {
						lastChar = grid[y][x]
						state = !state
					} else {
						if (lastChar == 'J' && grid[y][x] == 'L') ||
							(lastChar == 'L' && grid[y][x] == 'J') ||
							(lastChar == '7' && grid[y][x] == 'F') ||
							(lastChar == 'F' && grid[y][x] == '7') {
							state = !state
						}
						lastChar = '.'
					}

				}
			} else {
				lastChar = '.'
				if state {
					count++
				}
			}
		}
		state = false
	}

	return count
}
