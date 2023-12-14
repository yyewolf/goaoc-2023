package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

var HCols = make(map[int][][]int, 1000)
var HLines = make(map[int][][]int, 1000)

func doPartOne(input []byte) int {
	var lenGrid int

	var Os = make([][]int, 0, 1000)

	var x, y int
	for _, c := range input {
		if c == '\n' {
			lenGrid++
			x = 0
			y++
			continue
		}

		if c == 'O' {
			Os = append(Os, []int{y, x})
		}
		if c == '#' {
			HCols[x] = append(HCols[x], []int{y, x})
			HLines[y] = append(HLines[y], []int{y, x})
		}

		x++
	}

	// North
	// Calculate Os
	for i, o := range Os {
		// Find out where's the nearest H going up
		var nearestH int = -1
		for _, h := range HCols[o[1]] {
			if h[0] < o[0] {
				if h[0] > nearestH {
					nearestH = h[0]
				}
			}
		}

		Os[i][0] = nearestH + 1
	}

	// Move Os that are on the same box
	for _, o := range Os {
		// Count how many Os are on the same box
		var count = 0
		for j, o2 := range Os {
			if o[0] == o2[0] && o[1] == o2[1] {
				Os[j][0] += count
				count++
			}
		}
	}

	sum := 0
	for _, o := range Os {
		sum += lenGrid - o[0]
	}

	return sum
}

func doPartTwo(input []byte) int {
	var buf = bufio.NewReader(bytes.NewReader(input))

	var grid [][]byte

	var HCols = make(map[int][][]int, 1000)
	var HLines = make(map[int][][]int, 1000)

	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			break
		}

		grid = append(grid, line[:len(line)-1])
	}

	var Os = make([][]int, 0, 1000)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'O' {
				Os = append(Os, []int{i, j})

				// UNCOMMENT TO DEBUG :
				grid[i][j] = '.'

			}
			if grid[i][j] == '#' {
				HCols[j] = append(HCols[j], []int{i, j})
				HLines[i] = append(HLines[i], []int{i, j})
			}
		}
	}

	var loads = make([]int, 0, 100000)

	// Make all the 0 fall from gravity
	for iter := 0; iter < 1000000000; iter++ {
		for dir := 0; dir < 4; dir++ {
			for {
				var changed = false
				switch dir {
				case 0:
					// North
					// Recalculate Os
					for i, o := range Os {
						// Find out where's the nearest H going up
						var nearestH int = -1
						for _, h := range HCols[o[1]] {
							if h[0] < o[0] {
								if h[0] > nearestH {
									nearestH = h[0]
								}
							}
						}

						Os[i][0] = nearestH + 1
					}

					// Move Os that are on the same box
					for _, o := range Os {
						// Count how many Os are on the same box
						var count = 0
						for j, o2 := range Os {
							if o[0] == o2[0] && o[1] == o2[1] {
								Os[j][0] += count
								count++
							}
						}
					}
				case 1:
					// West

					for i, o := range Os {
						// Find out where's the nearest H going left
						var nearestH int = -1
						for _, h := range HLines[o[0]] {
							if h[1] < o[1] {
								if h[1] > nearestH {
									nearestH = h[1]
								}
							}
						}

						Os[i][1] = nearestH + 1
					}

					// Move Os that are on the same box
					for _, o := range Os {
						// Count how many Os are on the same box
						var count = 0
						for j, o2 := range Os {
							if o[0] == o2[0] && o[1] == o2[1] {
								Os[j][1] += count
								count++
							}
						}
					}
				case 2:
					// South

					for i, o := range Os {
						// Find out where's the nearest H going down
						var nearestH int = len(grid)
						for _, h := range HCols[o[1]] {
							if h[0] > o[0] {
								if h[0] < nearestH {
									nearestH = h[0]
								}
							}
						}

						Os[i][0] = nearestH - 1
					}

					// Move Os that are on the same box
					for _, o := range Os {
						// Count how many Os are on the same box
						var count = 0
						for j, o2 := range Os {
							if o[0] == o2[0] && o[1] == o2[1] {
								Os[j][0] -= count
								count++
							}
						}
					}
				case 3:
					// East

					for i, o := range Os {
						// Find out where's the nearest H going right
						var nearestH int = len(grid[0])
						for _, h := range HLines[o[0]] {
							if h[1] > o[1] {
								if h[1] < nearestH {
									nearestH = h[1]
								}
							}
						}

						Os[i][1] = nearestH - 1
					}

					// Move Os that are on the same box
					for _, o := range Os {
						// Count how many Os are on the same box
						var count = 0
						for j, o2 := range Os {
							if o[0] == o2[0] && o[1] == o2[1] {
								Os[j][1] -= count
								count++
							}
						}
					}
				}

				// Clear
				if !changed {
					break
				}
			}
		}

		load := 0
		for _, o := range Os {
			load += len(grid) - o[0]
		}

		loads = append(loads, load)

		// Check if we have a cycle of loads
		// This would mean that we have at least two times the same pattern :
		// 2 4 8 3  2 4 8 3
		for i := 0; i < len(loads); i++ {
			for patternSize := 2; patternSize < len(loads)-i; patternSize++ {
				if patternSize*2 > len(loads)-i {
					break
				}
				// Check if they're the same
				once := loads[i : i+patternSize]
				twice := loads[i+patternSize : i+patternSize*2]
				if fmt.Sprintf("%v", once) == fmt.Sprintf("%v", twice) {
					// Pattern found
					// fmt.Printf("Pattern found at %d, size %d\n", i, patternSize)
					// fmt.Printf("Pattern is %v\n", once)

					// Predict where we'll be at i = 1000000000

					return once[(1000000000-i-1)%patternSize]
				}
			}
		}
	}

	// for _, o := range Os {
	// 	grid[o[0]][o[1]] = 'O'
	// }

	// for i := 0; i < len(grid); i++ {
	// 	println(string(grid[i]))
	// }

	sum := 0
	for _, o := range Os {
		sum += len(grid) - o[0]
	}

	return sum
}
