package main

import (
	"iter"
	"math"
	"strconv"
	"strings"
)

type Equation struct {
	Result  int
	Members []int
}

var operators = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
	"||": func(a, b int) int {
		// concanate a and b using bit manipulation, we need to shift a to the left by 10^x, where x is the amount of digits in b
		// then we can just add them together
		i := int(math.Log10(float64(b)))

		return a*int(math.Pow10(i+1)) + b
	},
}

func getPermutationOfTwo(amountOfOperators int) iter.Seq[[]string] {
	possibilities := 1 << amountOfOperators
	return func(yield func([]string) bool) {
		for i := 0; i < possibilities; i++ {
			possibility := i
			// Get the operators
			op := make([]string, amountOfOperators)
			for j := 0; j < amountOfOperators; j++ {
				mod := possibility % 2
				switch mod {
				case 0:
					op[j] = "+"
				case 1:
					op[j] = "*"
				}
				possibility /= 2
			}
			if !yield(op) {
				return
			}
		}
	}
}

func (e *Equation) isSolvable() bool {
	// Check all possibilities, if one is correct, return true
	for operations := range getPermutationOfTwo(len(e.Members) - 1) {
		// Calculate the result from left to right
		result := e.Members[0]
		for j, operation := range operations {
			result = operators[operation](result, e.Members[j+1])
		}

		if result == e.Result {
			return true
		}
	}
	return false
}

func doPartOne(input string) int {
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

		if equation.isSolvable() {
			result += equation.Result
		}
	}

	return result // 1260333054159
}
