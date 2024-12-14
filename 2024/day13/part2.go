package main

import (
	"fmt"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var sum int

	for len(lines) > 0 {
		var game game
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &game.Ax, &game.Ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &game.Bx, &game.By)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &game.Gx, &game.Gy)
		lines = lines[3:]
		if len(lines) != 0 {
			lines = lines[1:]
		}

		game.Gx += 10000000000000
		game.Gy += 10000000000000

		sum += game.Solve()
	}

	return sum
}
