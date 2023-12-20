package main

import (
	"bufio"
	"bytes"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func checkGridMirrors(grid [][]byte) int {
	// To find the reflection in each pattern,
	// you need to find a perfect reflection across either a horizontal line
	// between two rows or across a vertical line between two columns.

	var doCheck = func(grid [][]byte) int {
		for i := 0; i < len(grid); i++ {
			// Check if current line is a mirror of the next line
			if i+1 < len(grid) && bytes.Equal(grid[i], grid[i+1]) {
				// Check if the reflection is perfect, each line is a mirror of the other going from the middle
				// If we hit a border, we can stop and it's a perfect reflection
				// If we hit a non-mirror, it's not a perfect reflection

				var isPerfect = true
				var k int
				for j := i + 1; j < len(grid); j++ {
					k++
					if i+1-k < 0 {
						break
					}
					if !bytes.Equal(grid[j], grid[i+1-k]) {
						isPerfect = false
						break
					}
				}

				if isPerfect {
					return i + 1
				}
			}
		}

		return 0
	}

	s := doCheck(grid)
	if s != 0 {
		return s * 100
	}

	// Transpose the grid
	var gridT [][]byte = make([][]byte, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		gridT[i] = make([]byte, len(grid))
		for j := 0; j < len(grid); j++ {
			gridT[i][j] = grid[j][i]
		}
	}

	s = doCheck(gridT)
	return s
}

func doPartOne(input []byte) int {
	// Search for rows mirror (repeating lines)
	var buf = bufio.NewReader(bytes.NewReader(input))

	var grid [][]byte = make([][]byte, 0, 100)

	sum := 0
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			break
		}

		if bytes.Equal(line, []byte{'\n'}) {
			s := checkGridMirrors(grid)
			// fmt.Println(s)
			sum += s
			grid = grid[:0]
		} else {
			grid = append(grid, line[:len(line)-1])
		}
	}
	sum += checkGridMirrors(grid)

	return sum
}

func checkGridMirrorsSmudge(grid [][]byte) int {
	// To find the reflection in each pattern,
	// you need to find a perfect reflection across either a horizontal line
	// between two rows or across a vertical line between two columns.

	var doCheck = func(grid [][]byte) int {
		for i := 0; i < len(grid); i++ {
			// Check if the reflection is perfect, each line is a mirror of the other going from the middle
			// If we hit a border, we can stop and it's a perfect reflection
			// If we hit a non-mirror, it's not a perfect reflection

			var smudge = false
			var isPerfect = true
			var k int
			for j := i + 1; j < len(grid); j++ {
				k++
				if i+1-k < 0 {
					break
				}
				if !bytes.Equal(grid[j], grid[i+1-k]) {

					// Check diff
					diff := 0
					if !smudge {
						for l := 0; l < len(grid[j]); l++ {
							if grid[j][l] != grid[i+1-k][l] {
								diff++
							}
						}
					}

					if diff == 1 {
						smudge = true
						continue
					}

					isPerfect = false
					break
				}
			}

			if isPerfect && smudge {
				return i + 1
			}
		}

		return 0
	}

	s := doCheck(grid)
	if s != 0 {
		return s * 100
	}

	// Transpose the grid
	gridT := make([][]byte, len(grid[0]))
	Q := make([]byte, len(grid)*len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		gridT[i] = Q[i*len(grid) : (i+1)*len(grid)]
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			gridT[i][j] = grid[j][i]
		}
	}

	s = doCheck(gridT)
	return s
}

func doPartTwo(input []byte) int {
	// Search for rows mirror (repeating lines)
	var buf = bufio.NewReader(bytes.NewReader(input))

	var grid [][]byte = make([][]byte, 0, 100)

	sum := 0
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			break
		}

		if bytes.Equal(line, []byte{'\n'}) {
			sum += checkGridMirrorsSmudge(grid)
			// fmt.Printf("Sum: %d\n", sum)
			grid = grid[:0]
		} else {
			grid = append(grid, line[:len(line)-1])
		}
	}
	sum += checkGridMirrorsSmudge(grid)

	// if sum != 37876 {
	// 	fmt.Println("Wrong answer")
	// }

	return sum
}
