package main

import (
	"bytes"
	"container/heap"
	"slices"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func dijkstra(grid [][]byte, maxConsecutive, movesNeededBeforeTurn int) int {
	m := len(grid)
	n := len(grid[0])
	startRight := state{row: 0, col: 0, dir: right, moves: 0}
	startDown := state{row: 0, col: 0, dir: down, moves: 0}
	pq := priorityQueue{
		&item{heatLoss: 0, state: startRight, index: 0},
		&item{heatLoss: 0, state: startDown, index: 1},
	}

	minCost := map[state]int{startRight: 0, startDown: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*item)
		if minCost[curr.state] < curr.heatLoss {
			continue
		}

		if curr.state.row == m-1 && curr.state.col == n-1 && curr.state.moves >= movesNeededBeforeTurn {
			return curr.heatLoss
		}

		for _, dir := range [4]direction{left, right, up, down} {
			if curr.state.moves == maxConsecutive && !slices.Contains(rotations[curr.state.dir], dir) || dir == reverse[curr.state.dir] {
				continue
			}

			ni, nj := curr.state.row+dir.row, curr.state.col+dir.col
			nextMoves := curr.state.moves

			if curr.state.moves < movesNeededBeforeTurn {
				if dir != curr.state.dir {
					continue
				}
				nextMoves += 1
			} else {
				if dir != curr.state.dir {
					nextMoves = 1
				} else {
					nextMoves = nextMoves%maxConsecutive + 1
				}
			}

			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}

			nextState := state{row: ni, col: nj, moves: nextMoves, dir: dir}
			nextHeatLoss := int(rune(grid[ni][nj]) - '0')
			if _, ok := minCost[nextState]; ok && minCost[nextState] <= curr.heatLoss+nextHeatLoss {
				continue
			}

			minCost[nextState] = curr.heatLoss + nextHeatLoss
			heap.Push(&pq, &item{heatLoss: curr.heatLoss + nextHeatLoss, state: nextState})
		}
	}

	return -1
}

func doPartOne(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))

	return dijkstra(grid, 3, 0)
}

func doPartTwo(input []byte) int {
	grid := bytes.Split(input, []byte("\n"))

	return dijkstra(grid, 10, 4)
}
