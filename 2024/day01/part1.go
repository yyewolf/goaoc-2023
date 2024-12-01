package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

//🎄 Running with real input for day 1 of year 2024 :
// 2367773
// 21271939

type List []int

// doPartOne solves the first part of the problem.
// The problem requires us to pair the smallest number from each list
func doPartOne(input string) int {
	var lists [2]List

	// In the input, lists are separated by three spaces
	// They are also represented vertically, so we can split by newlines
	for i, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, "   ")
		if len(numbers) != 2 {
			continue
		}

		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("Error parsing number on line %d (`%s`) : %v", i, line, err)
		}

		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("Error parsing number on line %d (`%s`) : %v", i, line, err)
		}

		lists[0] = append(lists[0], firstNumber)
		lists[1] = append(lists[1], secondNumber)
	}

	// Sort the lists
	for i := 0; i < len(lists); i++ {
		slices.Sort(lists[i])
	}

	// Since the lists are sorted, we can now simply sum the differences
	var sum int

	for i := 0; i < len(lists[0]); i++ {
		sum += max(lists[0][i], lists[1][i]) - min(lists[0][i], lists[1][i])
	}

	return sum
}
