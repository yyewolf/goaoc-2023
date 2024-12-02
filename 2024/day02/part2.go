package main

import (
	"strconv"
	"strings"
)

func doPartTwo(input string) int {
	var safeSteps int

	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		var steps []int
		for _, field := range fields {
			i, _ := strconv.Atoi(field)
			steps = append(steps, i)
		}

		if isSafe(steps) {
			safeSteps++
		} else {
			var newStep []int
			for i := 0; i < len(steps); i++ {
				newStep = newStep[:0]
				newStep = append(newStep, steps[:i]...)
				newStep = append(newStep, steps[i+1:]...)
				if isSafe(newStep) {
					safeSteps++
					break
				}
			}
		}
	}

	return safeSteps // 293
}
