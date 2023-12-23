package main

import (
	"bytes"
	"sort"
)

func simplifyGraph(graph map[int][][2]int) map[int][][2]int {
	simplifiedGraph := make(map[int][][2]int)
	var keys []int

	for node, neighbors := range graph {
		simplifiedGraph[node] = make([][2]int, len(neighbors))
		copy(simplifiedGraph[node], neighbors)
		keys = append(keys, node)
	}

	sort.Ints(keys)

	for _, node := range keys {
		neighbors := simplifiedGraph[node]
		if len(neighbors) == 2 {
			// If the node has only two neighbors, remove the node and connect the neighbors
			// to each other
			neighbor1 := neighbors[0][0]
			neighbor2 := neighbors[1][0]
			dist := neighbors[0][1] + neighbors[1][1]

			// Remove the node from the neighbors
			for i, neighbor := range simplifiedGraph[neighbor1] {
				if neighbor[0] == node {
					simplifiedGraph[neighbor1] = append(simplifiedGraph[neighbor1][:i], simplifiedGraph[neighbor1][i+1:]...)
					break
				}
			}

			for i, neighbor := range simplifiedGraph[neighbor2] {
				if neighbor[0] == node {
					simplifiedGraph[neighbor2] = append(simplifiedGraph[neighbor2][:i], simplifiedGraph[neighbor2][i+1:]...)
					break
				}
			}

			// Remove the node from the graph
			delete(simplifiedGraph, node)

			// Connect the neighbors
			simplifiedGraph[neighbor1] = append(simplifiedGraph[neighbor1], [2]int{neighbor2, dist})
			simplifiedGraph[neighbor2] = append(simplifiedGraph[neighbor2], [2]int{neighbor1, dist})
		}
	}

	return simplifiedGraph
}

func doPartTwo(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))

	var endX int = len(grid[0]) - 2
	var endY int = len(grid) - 1
	var width int = endX + 1
	var graph = make(map[int][][2]int)

	// Fill the graph
	for y := 0; y < endY+1; y++ {
		for x := 0; x < endX+1; x++ {
			if grid[y][x] == '#' {
				continue
			}

			for _, dir := range []string{"^", ">", "v", "<"} {
				newX, newY := getNeighborCoordinates(x, y, dir)

				if newX < 0 || newX > endX || newY < 0 || newY > endY || grid[newY][newX] == '#' {
					continue
				}

				graph[y*width+x] = append(graph[y*width+x], [2]int{newY*width + newX, 1})
			}
		}
	}

	graph = simplifyGraph(graph)

	// Perform DFS to find the longest path
	visited := make(map[int]bool)

	var dfs func(n [2]int) [][2]int
	dfs = func(n [2]int) [][2]int {
		node := n[0]
		dist := n[1]
		if visited[node] {
			return [][2]int{}
		}

		if node == endY*width+endX {
			return [][2]int{{node, dist}}
		}

		visited[node] = true
		maxPathLen := 0
		var maxPath [][2]int

		for _, neighbor := range graph[node] {
			path := dfs(neighbor)

			if len(path) == 0 || path[0][0] != endY*width+endX {
				continue
			}

			var pathSize int
			for _, n := range path {
				pathSize += n[1]
			}

			if pathSize > maxPathLen {
				maxPathLen = pathSize
				maxPath = path
			}
		}

		maxPath = append(maxPath, n)

		visited[node] = false
		return maxPath
	}

	p := dfs([2]int{1, 1})

	var pathSize int
	for _, n := range p {
		pathSize += n[1]
	}

	r := pathSize
	if r != 6379 {
		println("Wrong answer, expected 6379 but got", r)
	}
	return pathSize
}

func getNeighborCoordinates(x, y int, dir string) (int, int) {
	switch dir {
	case "^":
		return x, y - 1
	case ">":
		return x + 1, y
	case "v":
		return x, y + 1
	case "<":
		return x - 1, y
	}
	return x, y
}
