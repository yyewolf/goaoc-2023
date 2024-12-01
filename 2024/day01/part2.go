package main

import (
	"log"
	"strconv"
	"strings"
)

// This time, we want to calculate the similarity between the two lists
// This is done by counting how many times each number of list 1 appears in list 2 and multiplying this count by the number itself
// Parsing changes to make use of maps instead of slices
func doPartTwo(input string) int {
	// similarityMap will store the number of times each number appears in list 2, we'll deal about list 1 later
	var similarityMap = make(map[int]int)
	var list []int

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

		list = append(list, firstNumber)
		similarityMap[secondNumber]++
	}

	// For each number of list 1, we count how many times it appears in list 2
	var sum int
	for _, num := range list {
		sum += num * similarityMap[num]
	}

	return sum
}
