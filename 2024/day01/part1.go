package main

import (
	"slices"
	"strconv"
	"strings"
)

func doPartOne(input string) int {
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

	// Sort the lists
	slices.Sort(listOne)
	slices.Sort(listTwo)

	// Get the result of sum of matched numbers
	var sum int

	// Take out the min of the two lists, until both are empty
	for len(listOne) > 0 && len(listTwo) > 0 {
		minOne := listOne[0]
		minTwo := listTwo[0]

		sum += max(minOne, minTwo) - min(minOne, minTwo)

		listOne = listOne[1:]
		listTwo = listTwo[1:]
	}

	return sum
}
