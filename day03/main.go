package main

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	answer := doPartOne(input)
	fmt.Println(answer)

	answer = doPartTwo(input)
	fmt.Println(answer)
}

func isSymbol(c byte) bool {
	return c != '.' && !isNumber(c) && c != '\n'
}

func isStar(c byte) bool {
	return c == '*'
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func doPartOne(input []byte) int {
	width := 0
	r := bufio.NewReader(bytes.NewBuffer(input))
	sum := 0
	y := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if width == 0 {
			width = len(line) + 1
		}

		temp := 0
		hasAdjacency := false
		for x, c := range line {
			// If c is a number, check right for other number so that we can continue to read the number
			if !isNumber(c) {
				continue
			}

			if temp == 0 {
				// Check left, up-left, and down-left
				if x > 0 && isSymbol(line[x-1]) {
					hasAdjacency = true
				}
				if y > 0 && isSymbol(input[(y-1)*width+(x-1)]) {
					hasAdjacency = true
				}
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x-1]) {
					hasAdjacency = true
				}
			} else {
				// Only check up-left and down-left
				if y > 0 && isSymbol(input[(y-1)*width+(x-1)]) {
					hasAdjacency = true
				}
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x-1]) {
					hasAdjacency = true
				}
			}

			temp = temp*10 + int(c-'0')

			// If we're on the bottom right, only check up
			if x == width-2 && y == len(input)/width-1 {
				if y > 0 && isSymbol(input[(y-1)*width+x]) {
					hasAdjacency = true
				}

				// We're done with this one
				if hasAdjacency {
					sum += temp
				}
				temp = 0
				hasAdjacency = false
				continue
			}

			// If we're on the top right, check down only
			if x == width-2 && y == 0 {
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x]) {
					hasAdjacency = true
				}

				// We're done with this one
				if hasAdjacency {
					sum += temp
				}
				temp = 0
				hasAdjacency = false
				continue
			}

			// If we're on the last col, check up and down only
			if x == width-2 {
				if y > 0 && isSymbol(input[(y-1)*width+x]) {
					hasAdjacency = true
				}
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x]) {
					hasAdjacency = true
				}

				// We're done with this one
				if hasAdjacency {
					sum += temp
				}
				temp = 0
				hasAdjacency = false
				continue
			}

			// If we're on the last one, check right, up-right, and down-right and up and down
			if !isNumber(line[x+1]) {
				if x < width-2 && isSymbol(line[x+1]) {
					hasAdjacency = true
				}
				if y > 0 && isSymbol(input[(y-1)*width+x+1]) {
					hasAdjacency = true
				}
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x+1]) {
					hasAdjacency = true
				}
				if y > 0 && isSymbol(input[(y-1)*width+x]) {
					hasAdjacency = true
				}
				if y < len(input)/width-1 && isSymbol(input[(y+1)*width+x]) {
					hasAdjacency = true
				}

				// We're done with this one
				if hasAdjacency {
					sum += temp
				}
				temp = 0
				hasAdjacency = false
			}
		}
		y++
	}

	return sum
}

var connections = make([]int, 10)
var emptyConnections = make([]int, 10)
var gears = make(map[int][]int)

func doPartTwo(input []byte) int {
	width := 0
	r := bufio.NewReader(bytes.NewBuffer(input))

	y := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if width == 0 {
			width = len(line) + 1
		}

		// empty connections fast
		copy(connections, emptyConnections)
		temp := 0
		for x, c := range line {
			// If c is a number, check right for other number so that we can continue to read the number
			if !isNumber(c) {
				continue
			}

			if temp == 0 {
				// Check left, up-left, and down-left
				if x > 0 && isStar(line[x-1]) {
					connections = append(connections, y*width+x-1)
				}
				if y > 0 && isStar(input[(y-1)*width+(x-1)]) {
					connections = append(connections, (y-1)*width+x-1)
				}
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x-1]) {
					connections = append(connections, (y+1)*width+x-1)
				}
			} else {
				// Only check up-left and down-left
				if y > 0 && isStar(input[(y-1)*width+(x-1)]) {
					connections = append(connections, (y-1)*width+x-1)
				}
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x-1]) {
					connections = append(connections, (y+1)*width+x-1)
				}
			}

			temp = temp*10 + int(c-'0')

			// If we're on the bottom right, only check up
			if x == width-2 && y == len(input)/width-1 {
				if y > 0 && isStar(input[(y-1)*width+x]) {
					connections = append(connections, (y-1)*width+x)
				}

				// We're done with this one
				for _, k := range connections {
					gears[k] = append(gears[k], temp)
				}
				connections = connections[:0]
				temp = 0
				continue
			}

			// If we're on the top right, check down only
			if x == width-2 && y == 0 {
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
					connections = append(connections, (y+1)*width+x)
				}

				// We're done with this one
				for _, k := range connections {
					gears[k] = append(gears[k], temp)
				}
				connections = connections[:0]
				temp = 0
				continue
			}

			// If we're on the last col, check up and down only
			if x == width-2 {
				if y > 0 && isStar(input[(y-1)*width+x]) {
					connections = append(connections, (y-1)*width+x)
				}
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
					connections = append(connections, (y+1)*width+x)
				}

				// We're done with this one
				for _, k := range connections {
					gears[k] = append(gears[k], temp)
				}
				connections = connections[:0]
				temp = 0
				continue
			}

			// If we're on the last one, check right, up-right, and down-right and up and down
			if !isNumber(line[x+1]) {
				if x < width-2 && isStar(line[x+1]) {
					connections = append(connections, y*width+x+1)
				}
				if y > 0 && isStar(input[(y-1)*width+x+1]) {
					connections = append(connections, (y-1)*width+x+1)
				}
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x+1]) {
					connections = append(connections, (y+1)*width+x+1)
				}
				if y > 0 && isStar(input[(y-1)*width+x]) {
					connections = append(connections, (y-1)*width+x)
				}
				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
					connections = append(connections, (y+1)*width+x)
				}

				// We're done with this one
				for _, k := range connections {
					gears[k] = append(gears[k], temp)
				}
				connections = connections[:0]
				temp = 0
			}
		}
		y++
	}

	sum := 0
	for _, v := range gears {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}

	return sum
}

// func doPartTwoB(input []byte) int {
// 	width := 0
// 	y := 0
// 	var gears = make(map[int][]int)
// 	var connections = make([]int, 10)
// 	var emptyConnections = make([]int, 10)
// 	for i := 0; i < len(input); i++ {
// 		c := input[i]
// 		if c == '\n' {
// 			continue
// 		}
// 		if width == 0 {
// 			width = bytes.IndexByte(input, '\n') + 1
// 		}

// 		// empty connections fast
// 		copy(connections, emptyConnections)
// 		temp := 0
// 		for x, c := range line {
// 			// If c is a number, check right for other number so that we can continue to read the number
// 			if !isNumber(c) {
// 				continue
// 			}

// 			if temp == 0 {
// 				// Check left, up-left, and down-left
// 				if x > 0 && isStar(line[x-1]) {
// 					connections = append(connections, y*width+x-1)
// 				}
// 				if y > 0 && isStar(input[(y-1)*width+(x-1)]) {
// 					connections = append(connections, (y-1)*width+x-1)
// 				}
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x-1]) {
// 					connections = append(connections, (y+1)*width+x-1)
// 				}
// 			} else {
// 				// Only check up-left and down-left
// 				if y > 0 && isStar(input[(y-1)*width+(x-1)]) {
// 					connections = append(connections, (y-1)*width+x-1)
// 				}
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x-1]) {
// 					connections = append(connections, (y+1)*width+x-1)
// 				}
// 			}

// 			temp = temp*10 + int(c-'0')

// 			// If we're on the bottom right, only check up
// 			if x == width-2 && y == len(input)/width-1 {
// 				if y > 0 && isStar(input[(y-1)*width+x]) {
// 					connections = append(connections, (y-1)*width+x)
// 				}

// 				// We're done with this one
// 				for _, k := range connections {
// 					gears[k] = append(gears[k], temp)
// 				}
// 				copy(connections, emptyConnections)
// 				temp = 0
// 				continue
// 			}

// 			// If we're on the top right, check down only
// 			if x == width-2 && y == 0 {
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
// 					connections = append(connections, (y+1)*width+x)
// 				}

// 				// We're done with this one
// 				for _, k := range connections {
// 					gears[k] = append(gears[k], temp)
// 				}
// 				copy(connections, emptyConnections)
// 				temp = 0
// 				continue
// 			}

// 			// If we're on the last col, check up and down only
// 			if x == width-2 {
// 				if y > 0 && isStar(input[(y-1)*width+x]) {
// 					connections = append(connections, (y-1)*width+x)
// 				}
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
// 					connections = append(connections, (y+1)*width+x)
// 				}

// 				// We're done with this one
// 				for _, k := range connections {
// 					gears[k] = append(gears[k], temp)
// 				}
// 				copy(connections, emptyConnections)
// 				temp = 0
// 				continue
// 			}

// 			// If we're on the last one, check right, up-right, and down-right and up and down
// 			if !isNumber(line[x+1]) {
// 				if x < width-2 && isStar(line[x+1]) {
// 					connections = append(connections, y*width+x+1)
// 				}
// 				if y > 0 && isStar(input[(y-1)*width+x+1]) {
// 					connections = append(connections, (y-1)*width+x+1)
// 				}
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x+1]) {
// 					connections = append(connections, (y+1)*width+x+1)
// 				}
// 				if y > 0 && isStar(input[(y-1)*width+x]) {
// 					connections = append(connections, (y-1)*width+x)
// 				}
// 				if y < len(input)/width-1 && isStar(input[(y+1)*width+x]) {
// 					connections = append(connections, (y+1)*width+x)
// 				}

// 				// We're done with this one
// 				for _, k := range connections {
// 					gears[k] = append(gears[k], temp)
// 				}
// 				copy(connections, emptyConnections)
// 				temp = 0
// 			}
// 		}
// 		y++
// 	}

// 	sum := 0
// 	for _, v := range gears {
// 		if len(v) == 2 {
// 			sum += v[0] * v[1]
// 		}
// 	}

// 	return sum
// }
