package main

import "bytes"

func doPartTwo(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var height = len(lines)   // y axis
	var width = len(lines[0]) // x axis

	var graph = make(map[pos][]pos)
	var elevationZero []pos

	for y, line := range lines {
		for x, e := range line {
			currentPos := pos{x, y}

			if fastAtoi(e) == 0 {
				elevationZero = append(elevationZero, currentPos)
			}

			for _, neighbor := range neighbors {
				neighborPos := pos{x + neighbor[0], y + neighbor[1]}
				if neighborPos.x < 0 || neighborPos.x > width-1 || neighborPos.y < 0 || neighborPos.y > height-1 {
					continue
				}

				graph[currentPos] = append(graph[currentPos], neighborPos)
			}
		}
	}

	var sumScore int

	for _, start := range elevationZero {
		// For each start position, find out its score
		var score int

		var visited map[pos]bool = map[pos]bool{}
		var toVisit = []pos{start} // to visit is a FIFO queue (first in/first out), in order to do a BFS
		for len(toVisit) > 0 {
			atPos := toVisit[0]
			toVisit = toVisit[1:]

			if visited[atPos] {
				//continue // literally the only difference
			}

			visited[atPos] = true

			linkedTo := graph[atPos]
			for _, link := range linkedTo {
				// Only append unvisited links with an elevation diff of 1
				if !visited[link] && elevationDiff(lines, atPos, link) == 1 {
					toVisit = append(toVisit, link)
				}
			}

			if getElevation(lines, atPos) == 9 {
				score++
			}
		}

		sumScore += score
	}

	return sumScore
}
