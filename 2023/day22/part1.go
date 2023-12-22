package main

import (
	"bytes"
	"fmt"
	"sort"
)

type point struct {
	x int
	y int
	z int
}

type brick struct {
	id int
	p1 point
	p2 point
}

func lowestZ(b *brick) int {
	if b.p1.z < b.p2.z {
		return b.p1.z
	}

	return b.p2.z
}

func doPartOne(input []byte) int {
	var fallingBricks []*brick

	var occupied = make(map[point]int)
	var supportedBy = make(map[int]map[int]bool)

	for i, line := range bytes.Split(input, []byte("\n")) {
		var b brick

		fmt.Sscanf(string(line), "%d,%d,%d~%d,%d,%d", &b.p1.x, &b.p1.y, &b.p1.z, &b.p2.x, &b.p2.y, &b.p2.z)

		b.id = i

		fallingBricks = append(fallingBricks, &b)
	}

	// sort by lowest z
	sort.Slice(fallingBricks, func(i, j int) bool {
		return lowestZ(fallingBricks[i]) < lowestZ(fallingBricks[j])
	})

	// we can now make them fall in order
	for _, b := range fallingBricks {
		for z := lowestZ(b); z > -1; z-- {
			newP1 := point{b.p1.x, b.p1.y, z}
			newP2 := point{b.p2.x, b.p2.y, z}
			if b.p1.z < b.p2.z {
				diff := b.p2.z - b.p1.z
				newP1.z = z
				newP2.z = z + diff
			} else {
				diff := b.p1.z - b.p2.z
				newP2.z = z
				newP1.z = z + diff
			}

			// Check if any points are occupied
			var able = true
			supportedBy[b.id] = make(map[int]bool)
			for x := newP1.x; x <= newP2.x; x++ {
				for y := newP1.y; y <= newP2.y; y++ {
					if _, ok := occupied[point{x, y, z}]; ok {
						// We can't move this brick
						able = false
						supportedBy[b.id][occupied[point{x, y, z}]] = true
					}
				}
			}

			if able && z != 0 {
				continue
			}

			if !able {
				newP1.z += 1
				newP2.z += 1
			}

			// Move the brick
			b.p1 = newP1
			b.p2 = newP2

			// mark as occupied
			for x := newP1.x; x <= newP2.x; x++ {
				for y := newP1.y; y <= newP2.y; y++ {
					for _z := newP1.z; _z <= newP2.z; _z++ {
						occupied[point{x, y, _z}] = b.id
					}
				}
			}

			break
		}
	}

	// Invert map to count which bricks support which
	var supporting = make(map[int][]int)

	for k, v := range supportedBy {
		for k2 := range v {
			supporting[k2] = append(supporting[k2], k)
		}
	}

	// Check which bricks can be deleted without collapsing the tower
	count := 0

	for _, b := range fallingBricks {
		v, ok := supporting[b.id]
		// Not supporting any bricks, can be deleted
		if !ok {
			count++
			continue
		}

		// Check if all bricks supported by this brick are supported by other bricks
		able := true
		for _, id := range v {
			if len(supportedBy[id]) == 1 {
				able = false
				break
			}
		}

		if able {
			count++
		}
	}

	return count
}
