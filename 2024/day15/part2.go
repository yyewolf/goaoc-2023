package main

import (
	"bytes"
	"fmt"
	"slices"
)

type cell struct {
	pos      [2]int
	value    byte
	linkedTo *cell
}

func (c *cell) String() string {
	return fmt.Sprintf("Cell{pos: %v, value: '%s', linkedTo: %p}", c.pos, string(c.value), c.linkedTo)
}

type warehouse struct {
	innerMap map[[2]int]*cell
	width    int
	height   int

	moves []byte
}

func parseAndWiden(input []byte) warehouse {
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[:len(lines)-1]
	var height = 0
	var width = len(lines[0])

	// Parsing the input
	var initialWarehouse map[[2]int]byte = make(map[[2]int]byte)
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
				initialWarehouse[[2]int{x, y}] = c
			}
		}
		if state == 1 {
			moves = append(moves, line...)
		}
	}

	// Widen the map, according to the new rules
	var newWarehouse map[[2]int]byte = make(map[[2]int]byte)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if initialWarehouse[[2]int{x, y}] == '#' {
				newWarehouse[[2]int{x * 2, y}] = '#'
				newWarehouse[[2]int{x*2 + 1, y}] = '#'
			} else if initialWarehouse[[2]int{x, y}] == 'O' {
				newWarehouse[[2]int{x * 2, y}] = '['
				newWarehouse[[2]int{x*2 + 1, y}] = ']'
			} else if initialWarehouse[[2]int{x, y}] == '.' {
				newWarehouse[[2]int{x * 2, y}] = '.'
				newWarehouse[[2]int{x*2 + 1, y}] = '.'
			} else if initialWarehouse[[2]int{x, y}] == '@' {
				newWarehouse[[2]int{x * 2, y}] = '@'
				newWarehouse[[2]int{x*2 + 1, y}] = '.'
			}
		}
	}

	var output warehouse = warehouse{innerMap: make(map[[2]int]*cell)}
	// Init warehouse
	for pos, val := range newWarehouse {
		output.innerMap[pos] = &cell{pos: pos, value: val}
	}
	// Do links for boxes
	for pos, val := range newWarehouse {
		if val == '[' {
			output.innerMap[pos].linkedTo = output.innerMap[[2]int{pos[0] + 1, pos[1]}]
		}
		if val == ']' {
			output.innerMap[pos].linkedTo = output.innerMap[[2]int{pos[0] - 1, pos[1]}]
		}
	}

	output.width = width * 2
	output.height = height
	output.moves = moves
	return output
}

func doPartTwo(input []byte) int {
	warehouse := parseAndWiden(input)

	// Find the robot
	var robotX, robotY int
	for y := 0; y < warehouse.height; y++ {
		for x := 0; x < warehouse.width; x++ {
			if warehouse.innerMap[[2]int{x, y}].value == '@' {
				robotX = x
				robotY = y
			}
		}
	}

	for len(warehouse.moves) > 0 {
		currentMove := warehouse.moves[0]
		warehouse.moves = warehouse.moves[1:]

		nextRobotX := robotX + availableMoves[currentMove][0]
		nextRobotY := robotY + availableMoves[currentMove][1]

		if warehouse.innerMap[[2]int{nextRobotX, nextRobotY}].value == '#' {
			continue
		}

		var hitsWall bool
		if warehouse.innerMap[[2]int{nextRobotX, nextRobotY}].linkedTo != nil {
			hitsWall = warehouse.moveBox(warehouse.innerMap[[2]int{nextRobotX, nextRobotY}], currentMove)
		}

		if hitsWall {
			continue
		}

		warehouse.innerMap[[2]int{robotX, robotY}].value = '.'
		robotX = nextRobotX
		robotY = nextRobotY
		warehouse.innerMap[[2]int{robotX, robotY}].value = '@'
	}

	var sum int
	for pos, val := range warehouse.innerMap {
		if val.value == '[' {
			sum += pos[0] + pos[1]*100
		}
	}

	return sum
}

func (w warehouse) moveBox(potentialBox *cell, direction byte) (hitsWall bool) {
	if potentialBox.linkedTo == nil {
		// meaning this is not a box
		return
	}

	// We will need to check all of those cells
	var stack []*cell = []*cell{potentialBox, potentialBox.linkedTo}

	var seen = make(map[[2]int]bool)

	// Find the last box in the direction of the move
	var at = 0
	for at < len(stack) {
		box := stack[at]
		seen[box.pos] = true
		nextPos := [2]int{box.pos[0] + availableMoves[direction][0], box.pos[1] + availableMoves[direction][1]}
		nextBox := w.innerMap[nextPos]

		if nextBox.linkedTo == nil {
			if nextBox.value == '#' {
				hitsWall = true
			}
			// meaning this is not a box
			at++
			continue
		}

		if seen[nextBox.pos] {
			// We already saw this box
			at++
			continue
		}

		stack = append(stack, nextBox, nextBox.linkedTo)
		at++
	}

	if hitsWall {
		return
	}

	// Move the boxes
	for len(stack) > 0 {
		box := stack[len(stack)-1]
		linkedTo := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		nextBox := w.innerMap[[2]int{box.pos[0] + availableMoves[direction][0], box.pos[1] + availableMoves[direction][1]}]
		nextLinkedTo := w.innerMap[[2]int{linkedTo.pos[0] + availableMoves[direction][0], linkedTo.pos[1] + availableMoves[direction][1]}]

		nextBoxCell := cell{pos: nextBox.pos, value: box.value, linkedTo: nextLinkedTo}
		nextLinkedToCell := cell{pos: nextLinkedTo.pos, value: linkedTo.value, linkedTo: &nextBoxCell}
		nextBoxCell.linkedTo = &nextLinkedToCell

		boxCell := cell{pos: box.pos, value: '.', linkedTo: nil}
		linkedToCell := cell{pos: linkedTo.pos, value: '.', linkedTo: nil}

		w.innerMap[boxCell.pos] = &boxCell
		w.innerMap[linkedToCell.pos] = &linkedToCell
		w.innerMap[nextBoxCell.pos] = &nextBoxCell
		w.innerMap[nextLinkedToCell.pos] = &nextLinkedToCell
	}
	return
}

// func (w warehouse) draw() byte {
// 	for y := 0; y < w.height; y++ {
// 		for x := 0; x < w.width; x++ {
// 			print(string(w.innerMap[[2]int{x, y}].value))
// 		}
// 		println()
// 	}
// 	return 0
// }
