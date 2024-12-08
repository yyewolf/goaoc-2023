package main

import (
	"bytes"
)

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var height int = len(lines)
	var width int = len(lines[0])

	var antennas = make(map[byte][][2]int)

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], [2]int{x, y})
		}
	}

	var antiNodes = make(map[[2]int]bool)
	// For each pair of antennas of the same type, we check if the antinodes are in bounds,
	// for a point A and a point B, one antinode would be the point C such that AC=2AB, the other antinode would be the point D such that BD=2AB
	for _, antenna := range antennas {
		for i, a := range antenna {
			for j := i + 1; j < len(antenna); j++ {
				b := antenna[j]
				if i == j {
					continue
				}
				// We calculate the vector AB
				ab := [2]int{b[0] - a[0], b[1] - a[1]}
				// We calculate the antinode C
				c := [2]int{a[0] + 2*ab[0], a[1] + 2*ab[1]}
				// We calculate the antinode D
				d := [2]int{b[0] - 2*ab[0], b[1] - 2*ab[1]}

				// We check if the antinode C is in bounds
				if c[0] >= 0 && c[0] < width && c[1] >= 0 && c[1] < height {
					antiNodes[c] = true
				}
				// We check if the antinode D is in bounds
				if d[0] >= 0 && d[0] < width && d[1] >= 0 && d[1] < height {
					antiNodes[d] = true
				}
			}
		}
	}

	// Draw the map with # for antinodes, . for empty space and the antennas
	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		if _, ok := antiNodes[[2]int{x, y}]; ok {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(string(lines[y][x]))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return len(antiNodes)
}
