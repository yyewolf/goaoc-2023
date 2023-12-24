package main

import (
	"bytes"
	"fmt"
)

type node struct {
	px, py, pz int
	vx, vy, vz int
}

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))

	// format: px py pz @ vx vy vz
	var nodes []*node
	for _, line := range lines {
		var n node
		_, err := fmt.Sscanf(string(line), "%d, %d, %d @ %d, %d, %d", &n.px, &n.py, &n.pz, &n.vx, &n.vy, &n.vz)
		if err != nil {
			panic(err)
		}

		nodes = append(nodes, &n)
	}

	s := 0

	var testAreaMin = 200000000000000
	var testAreaMax = 400000000000000

	for i, n := range nodes {
		// Convert node to
		// y = ax + b
		// a = vy / vx
		// b = py - a * px

		a := float64(n.vy) / float64(n.vx)
		b := float64(n.py) - a*float64(n.px)

		// Check if it will intersect with any other node
		// Only mind the x and y axis
		// Test if they collide within testAreaMin and testAreaMax
		for j := i + 1; j < len(nodes); j++ {
			n2 := nodes[j]
			if n == n2 {
				continue
			}

			// Convert node to
			// y = ax + b
			// a = vy / vx
			// b = py - a * px

			a2 := float64(n2.vy) / float64(n2.vx)
			b2 := float64(n2.py) - a2*float64(n2.px)

			// Check if they intersect
			// y = ax + b
			// y = a2x + b2
			// ax + b = a2x + b2

			// x = (b2 - b) / (a - a2)
			x := (b2 - b) / (a - a2)
			y := a*x + b

			t := (x - float64(n.px)) / float64(n.vx)
			t2 := (x - float64(n2.px)) / float64(n2.vx)

			if x > float64(testAreaMin) && x < float64(testAreaMax) && y > float64(testAreaMin) && y < float64(testAreaMax) && t > 0 && t2 > 0 {
				s++
			}
		}
	}

	return s
}
