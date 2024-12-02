package main

import (
	"strconv"
	"strings"
)

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func isSafe(steps []int) bool {
	var prevStep = steps[0]
	var isAscending, isDescending, isInRange = true, true, true

	for _, step := range steps[1:] {
		diff := abs(step - prevStep)
		if step > prevStep {
			isDescending = false
		} else if step < prevStep {
			isAscending = false
		}
		if diff > 3 || diff == 0 {
			isInRange = false
		}
		prevStep = step
	}

	return isInRange && (isAscending || isDescending)
}

func doPartOne(input string) int {
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
		}
	}

	return safeSteps // 224
}
