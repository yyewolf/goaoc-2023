package main

import (
	"fmt"
	"strings"
)

type robot struct {
	posX, posY int
	velX, velY int
}

var width = 101
var height = 103

func mod(a, b int) int {
	return (a%b + b) % b
}

func (r *robot) step(n int) {
	// Make the robot move n steps, if he is at position (x, y) and has velocity (vx, vy), then
	// after n steps he will be at position (x + n*vx, y + n*vy)
	// if he goes out of the grid, then he wraps around
	r.posX = mod(r.posX+n*r.velX, width)
	r.posY = mod(r.posY+n*r.velY, height)
}

func doPartOne(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var robots []robot
	for len(lines) > 0 {
		var robot robot
		fmt.Sscanf(lines[0], "p=%d,%d v=%d,%d", &robot.posX, &robot.posY, &robot.velX, &robot.velY)
		robots = append(robots, robot)
		lines = lines[1:]
	}

	var positions [103][101]int

	for _, robot := range robots {
		robot.step(100)
		positions[robot.posY][robot.posX]++
	}

	// count robots in each quadrant, not including the center lines
	var numbers [4]int
	for i := 0; i < height/2; i++ {
		for j := 0; j < width/2; j++ {
			// top left
			numbers[0] += positions[i][j]
			// top right
			numbers[1] += positions[i][j+width/2+1]
			// bottom left
			numbers[2] += positions[i+height/2+1][j]
			// bottom right
			numbers[3] += positions[i+height/2+1][j+width/2+1]
		}
	}

	var mul = 1
	for i := 0; i < 4; i++ {
		mul *= numbers[i]
	}

	return mul
}
