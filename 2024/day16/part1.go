package main

import (
	"bytes"
	"math"
)

// ###############
// #.......#....E#
// #.#.###.#.###.#
// #.....#.#...#.#
// #.###.#####.#.#
// #.#.#.......#.#
// #.#.#####.###.#
// #...........#.#
// ###.#.#####.#.#
// #...#.....#.#.#
// #.#.#.###.#.#.#
// #.....#...#.#.#
// #.###.#.#.#.#.#
// #S..#.....#...#
// ###############
//
// Turning 90 degrees to the right cost 1000
// Moving cost 1
// We start at S pointing to the right

type dir int

const (
	up dir = iota
	right
	down
	left
)

var rotationCosts = map[[2]dir]int{
	{up, up}:    0,
	{up, right}: 1000,
	{up, down}:  2000,
	{up, left}:  1000,

	{right, up}:    1000,
	{right, right}: 0,
	{right, down}:  1000,
	{right, left}:  2000,

	{down, up}:    2000,
	{down, right}: 1000,
	{down, down}:  0,
	{down, left}:  1000,

	{left, up}:    1000,
	{left, right}: 2000,
	{left, down}:  1000,
	{left, left}:  0,
}

type pos struct {
	x, y int
	dir  dir
}

type move struct {
	p       pos
	weight  int
	isFinal bool
}

type graph map[pos][]move

var neighbors = []pos{
	{0, -1, up},
	{1, 0, right},
	{0, 1, down},
	{-1, 0, left},
}

func doPartOne(input []byte) int {
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

	var toVisit = []move{{start, 0, false}}
	var visited = make(map[pos]move)
	var minWeight = math.MaxInt32
	for len(toVisit) > 0 {
		var current = toVisit[0]
		toVisit = toVisit[1:]

		if current.isFinal {
			if current.weight < minWeight {
				minWeight = current.weight
			}
			continue
		}

		if val, found := visited[current.p]; val.weight < current.weight && found {
			continue
		}

		visited[current.p] = current

		// Add all "." according to the current direction
		for _, neighbor := range neighbors {
			if lines[current.p.y+neighbor.y][current.p.x+neighbor.x] == '.' || lines[current.p.y+neighbor.y][current.p.x+neighbor.x] == 'E' {
				weight := 1 + rotationCosts[[2]dir{current.p.dir, neighbor.dir}]
				toVisit = append(toVisit, move{pos{current.p.x + neighbor.x, current.p.y + neighbor.y, neighbor.dir}, current.weight + weight, lines[current.p.y+neighbor.y][current.p.x+neighbor.x] == 'E'})
			}
		}
	}

	return minWeight
}
