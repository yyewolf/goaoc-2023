package main

import (
	"bytes"
	"slices"
)

// func drawWarehouse(warehouse map[[2]int]byte, width, height int) {
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			print(string(warehouse[[2]int{x, y}]))
// 		}
// 		println()
// 	}
// }

var availableMoves map[byte][2]int = map[byte][2]int{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func doPartOne(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]

	var height = 0
	var width = len(lines[0])

	// Parsing the input
	var warehouse map[[2]int]byte = make(map[[2]int]byte)
	var moves []byte
	var state = 0
	for y, line := range lines {
		if slices.Equal(line, []byte("")) {
			height = y
			state++
			continue
		}
		if state == 0 {
			for x, c := range line {
				warehouse[[2]int{x, y}] = c
			}
		}
		if state == 1 {
			moves = append(moves, line...)
		}
	}

	// Find the robot
	var robotX, robotY int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if warehouse[[2]int{x, y}] == '@' {
				robotX = x
				robotY = y
			}
		}
	}

	for len(moves) > 0 {
		// Move the robot
		// If the robot meets a box (O), move it if possible :
		//   Go to the last box in the direction of the move, check if there is a free space in the direction of the move and move all the boxes in the direction of the move
		//  If there is no free space, stop the robot
		// If the robot meets a wall (#), stop the robot
		currentMove := moves[0]
		moves = moves[1:]

		nextRobotX := robotX + availableMoves[currentMove][0]
		nextRobotY := robotY + availableMoves[currentMove][1]

		if warehouse[[2]int{nextRobotX, nextRobotY}] == '#' {
			continue
		}

		if warehouse[[2]int{nextRobotX, nextRobotY}] == 'O' {
			// Find the last box in the direction of the move
			lastBoxX := nextRobotX
			lastBoxY := nextRobotY
			endsWithWall := false

			var boxes = [][2]int{{lastBoxX, lastBoxY}}

			for {
				nextLastBoxX := lastBoxX + availableMoves[currentMove][0]
				nextLastBoxY := lastBoxY + availableMoves[currentMove][1]

				if warehouse[[2]int{nextLastBoxX, nextLastBoxY}] == '#' {
					endsWithWall = true
					break
				} else if warehouse[[2]int{nextLastBoxX, nextLastBoxY}] == 'O' {
					lastBoxX = nextLastBoxX
					lastBoxY = nextLastBoxY
					boxes = append(boxes, [2]int{nextLastBoxX, nextLastBoxY})
				} else if warehouse[[2]int{nextLastBoxX, nextLastBoxY}] == '.' {
					break
				}
			}

			if endsWithWall {
				continue
			}

			// Move the boxes, doing it efficiently is putting the first box, after the last box, in the direction of the move
			firstBox := boxes[0]

			warehouse[firstBox] = '.'
			warehouse[[2]int{lastBoxX + availableMoves[currentMove][0], lastBoxY + availableMoves[currentMove][1]}] = 'O'
		}

		warehouse[[2]int{robotX, robotY}] = '.'
		robotX = nextRobotX
		robotY = nextRobotY
		warehouse[[2]int{robotX, robotY}] = '@'
	}

	// drawWarehouse(warehouse, width, height)

	// Locate all the boxes according to the GPS coordinates
	// Its equal to 100 times its distance from the top edge of the warehouse
	// plus its distance from the left edge of the warehouse
	var sum int
	for pos, val := range warehouse {
		if val == 'O' {
			sum += pos[0] + pos[1]*100
		}
	}

	return sum
}
