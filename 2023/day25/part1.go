package main

import (
	"bytes"
)

var graph = make([][]int, 1644825)

func toInt(b []byte) int {
	var n int
	for _, c := range b {
		n = n<<8 + int(c-'a')
	}
	return n
}

func doPartOne(input []byte) int {
	// replace : by space
	input = bytes.Replace(input, []byte(":"), []byte(" "), -1)
	lines := bytes.Split(input, []byte("\n"))

	clear(graph)
	lenGraph := 0
	// We start by a random node (first in the map is not guaranteed to be the same each time)
	var start int = -1

	for _, line := range lines {
		parts := bytes.Split(line, []byte(" "))
		p0 := toInt(parts[0])
		if start == -1 {
			start = p0
		}
		for _, part := range parts[2:] {
			p := toInt(part)
			graph[p0] = append(graph[p0], p)
			graph[p] = append(graph[p], p0)

			if len(graph[p0]) == 1 {
				lenGraph++
			}
			if len(graph[p]) == 1 {
				lenGraph++
			}

			// graph[string(parts[0])] = append(graph[string(parts[0])], string(part))
			// graph[string(part)] = append(graph[string(part)], string(parts[0]))
		}
	}

	// Component is a subgraph of the graph
	// We need to create it by adding nodes one by one
	var subGraph1 = make(map[int]bool)
	subGraph1[start] = true

	var middleEdges = make(map[[2]int]bool)
	for _, end := range graph[start] {
		middleEdges[[2]int{start, end}] = true
	}

	// We stop at 3 because the two subgraphs will be connected by 3 edges
	for len(middleEdges) > 3 {
		// Find the node with the minimum number of edges to the subgraph
		// As such, we will not travel to the other subgraph
		var nextNode int
		var min = 999999999
		for node := range middleEdges {
			var count = 0
			for _, end := range graph[node[1]] {
				if subGraph1[end] {
					count--
				} else {
					count++
				}
			}
			if count < min {
				min = count
				nextNode = node[1]
			}
		}

		// Add the node to the subgraph
		subGraph1[nextNode] = true

		// If the node is connected to the other subgraph, remove the edge
		// Otherwise, add the edge to the middleEdges
		for _, end := range graph[nextNode] {
			if subGraph1[end] {
				delete(middleEdges, [2]int{end, nextNode})
			} else {
				middleEdges[[2]int{nextNode, end}] = true
			}
		}
	}

	subGraph1Size := len(subGraph1)
	subGraph2Size := lenGraph - subGraph1Size

	return subGraph1Size * subGraph2Size
}
