package main

import (
	"bytes"
	"sort"
)

func simplifyGraph(graph map[int][][2]int) {
	var keys []int

	for node := range graph {
		keys = append(keys, node)
	}

	sort.Ints(keys)

	for _, node := range keys {
		neighbors := graph[node]
		if len(neighbors) == 2 {
			// If the node has only two neighbors, remove the node and connect the neighbors
			// to each other
			neighbor1 := neighbors[0][0]
			neighbor2 := neighbors[1][0]
			dist := neighbors[0][1] + neighbors[1][1]

			// Remove the node from the neighbors
			for i, neighbor := range graph[neighbor1] {
				if neighbor[0] == node {
					graph[neighbor1] = append(graph[neighbor1][:i], graph[neighbor1][i+1:]...)
					break
				}
			}

			for i, neighbor := range graph[neighbor2] {
				if neighbor[0] == node {
					graph[neighbor2] = append(graph[neighbor2][:i], graph[neighbor2][i+1:]...)
					break
				}
			}

			// Remove the node from the graph
			delete(graph, node)

			// Connect the neighbors
			graph[neighbor1] = append(graph[neighbor1], [2]int{neighbor2, dist})
			graph[neighbor2] = append(graph[neighbor2], [2]int{neighbor1, dist})
		}
	}
}

func doPartTwo(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))

	var endX int = len(grid[0]) - 2
	var endY int = len(grid) - 1
	var width int = endX + 1
	var graph = make(map[int][][2]int, endX*endY)

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

	simplifyGraph(graph)

	// Convert graph to plantuml
	// fmt.Println("@startuml")
	// fmt.Println("digraph G {")
	// for node, neighbors := range graph {
	// 	for _, neighbor := range neighbors {
	// 		fmt.Printf("%d -> %d [label=\"%d\"]\n", node, neighbor[0], neighbor[1])
	// 	}
	// }
	// fmt.Println("}")

	// Perform DFS to find the longest path
	visited := make(map[int]bool, endX*endY)

	var dfs func(n [2]int) int
	dfs = func(n [2]int) int {
		node := n[0]
		dist := n[1]
		if visited[node] {
			return 0
		}

		if node == endY*width+endX {
			return dist
		}

		visited[node] = true
		maxPathLen := 0

		for _, neighbor := range graph[node] {
			path := dfs(neighbor)

			if path == 0 {
				continue
			}

			path += dist

			if path > maxPathLen {
				maxPathLen = path
			}
		}

		visited[node] = false

		return maxPathLen
	}

	r := dfs([2]int{1, 1})

	if r != 6379 {
		println("Wrong answer, expected 6379 but got", r)
	}
	return r
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
