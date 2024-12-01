package main

import (
	"strconv"
	"strings"
)

func doPartTwo(input string) int {
	var listOne []int
	var listTwo []int

	// Split input by `   ` and `\n`
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		splt := strings.Split(line, "   ")

		i, _ := strconv.Atoi(splt[0])
		listOne = append(listOne, i)

		i, _ = strconv.Atoi(splt[1])
		listTwo = append(listTwo, i)
	}

	var sum int

	// For each number of list 1, we count how many times it appears in list 2
	for _, num := range listOne {
		count := 0
		for _, num2 := range listTwo {
			if num == num2 {
				count++
			}
		}

		sum += num * count
	}

	return sum
}
