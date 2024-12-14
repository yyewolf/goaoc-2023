package main

import (
	"fmt"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var robots []robot
	for len(lines) > 0 {
		var robot robot
		fmt.Sscanf(lines[0], "p=%d,%d v=%d,%d", &robot.posX, &robot.posY, &robot.velX, &robot.velY)
		robots = append(robots, robot)
		lines = lines[1:]
	}

	var foundAt int
	var originalRobots = make([]robot, len(robots))
	copy(originalRobots, robots)

	for step := 1; step <= 100000; step++ {
		var positions [103][101]int
		copy(robots, originalRobots)
		for _, robot := range robots {
			robot.step(step)
			positions[robot.posY][robot.posX]++
		}

		// check if there's a line of 8 or more robots in a row
		var found bool
		for i := 0; i < height; i++ {
			var line int = 0
			for j := 0; j < width; j++ {
				if positions[i][j] > 0 {
					line++
				} else {
					line = 0
				}
				if line >= 8 {
					foundAt = step
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	return foundAt
}
