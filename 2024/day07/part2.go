package main

import (
	"strconv"
	"strings"
)

func (e *Equation) isSolvableTwo() bool {
	amountOfOperators := len(e.Members) - 1
	// 3 pow amountOfOperators
	possibilities := 1
	for i := 0; i < amountOfOperators; i++ {
		possibilities *= 3
	}
	// Check all possibilities, if one is correct, return true
	for i := 0; i < possibilities; i++ {
		// Get the current possibility
		possibility := i
		// Get the operators
		op := make([]string, amountOfOperators)
		for j := 0; j < amountOfOperators; j++ {
			mod := possibility % 3
			switch mod {
			case 0:
				op[j] = "+"
			case 1:
				op[j] = "*"
			case 2:
				op[j] = "||"
			}
			possibility /= 3
		}

		// Calculate the result from left to right
		result := e.Members[0]
		for j := 0; j < amountOfOperators; j++ {
			result = operators[op[j]](result, e.Members[j+1])
		}

		if result == e.Result {
			return true
		}
	}
	return false
}

func doPartTwo(input string) int {
	var result int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		equation := Equation{
			Result:  0,
			Members: []int{},
		}

		// result is before the first :
		equation.Result, _ = strconv.Atoi(strings.Split(line, ":")[0])

		// members are after the first :
		for _, member := range strings.Split(strings.Split(line, ": ")[1], " ") {
			memberInt, _ := strconv.Atoi(member)
			equation.Members = append(equation.Members, memberInt)
		}

		if equation.isSolvableTwo() {
			result += equation.Result
		}
	}

	return result // 162042343638683
}
