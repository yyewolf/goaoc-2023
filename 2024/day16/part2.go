package main

import (
	"bytes"
	"math"
)

type moveTwo struct {
	p       pos
	weight  int
	isFinal bool
	path    []pos
}

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var start pos
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = pos{x, y, right}
			}
		}
	}

	var toVisit = []moveTwo{{start, 0, false, []pos{start}}}
	var visited = make(map[pos]int)
	var minWeight = math.MaxInt32
	var paths [][]pos
	for len(toVisit) > 0 {
		var current = toVisit[0]
		toVisit = toVisit[1:]

		if current.isFinal {
			if current.weight < minWeight {
				minWeight = current.weight
				paths = [][]pos{current.path} // Reset paths with new minWeight
			} else if current.weight == minWeight {
				paths = append(paths, current.path)
			}
			continue
		}

		// Skip if the current weight is greater than the minimum weight found for this position
		if val, found := visited[current.p]; found && val < current.weight {
			continue
		}

		visited[current.p] = current.weight

		// Add all neighbors
		for _, neighbor := range neighbors {
			newX := current.p.x + neighbor.x
			newY := current.p.y + neighbor.y
			if newX >= 0 && newX < len(lines[0]) && newY >= 0 && newY < len(lines) {
				cell := lines[newY][newX]
				if cell == '.' || cell == 'E' {
					weight := 1 + rotationCosts[[2]dir{current.p.dir, neighbor.dir}]
					newPath := append([]pos{}, current.path...)
					newPath = append(newPath, pos{newX, newY, neighbor.dir})
					toVisit = append(toVisit, moveTwo{
						p:       pos{newX, newY, neighbor.dir},
						weight:  current.weight + weight,
						isFinal: cell == 'E',
						path:    newPath,
					})
				}
			}
		}
	}

	var tiles = make(map[[2]int]bool)

	for _, path := range paths {
		for _, p := range path {
			tiles[[2]int{p.x, p.y}] = true
		}
	}

	return len(tiles)
}
