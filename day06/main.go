package main

import (
	"math"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func doPartOne(input []byte) int {

	// Times on the first lije
	var races [][2]int
	var tempT int
	var tempD int

	lineLength := len(input) / 2

	var pos int
	for {
		c1 := input[pos]
		c2 := input[pos+lineLength]

		// If c1 is a space or num and c2 is a space or num
		if (c1 == ' ' || c1 >= '0' && c1 <= '9') && (c2 == ' ' || c2 >= '0' && c2 <= '9') && (c1 != ' ' || c2 != ' ') {
			if c1 != ' ' {
				tempT = tempT*10 + int(c1-'0')
			}
			if c2 != ' ' {
				tempD = tempD*10 + int(c2-'0')
			}
		} else if c1 == ' ' && c2 == ' ' || c1 == '\n' && c2 == '\n' {
			if tempD != 0 || tempT != 0 {
				races = append(races, [2]int{tempT, tempD})
			}
			tempT = 0
			tempD = 0
			if c1 == '\n' && c2 == '\n' {
				break
			}
		}
		pos++
	}

	var sum int = 1

	// For each race (time and record distance)
	for _, race := range races {
		time := race[0]
		record := race[1]

		var lowerBounds = [2]int{0, time}
		for {
			holdFor := (lowerBounds[0] + lowerBounds[1]) / 2
			distance := holdFor * (time - holdFor)
			if distance > record {
				lowerBounds[1] = holdFor
			} else {
				lowerBounds[0] = holdFor
			}
			if lowerBounds[1]-lowerBounds[0] <= 1 {
				break
			}
		}

		// find upper bound that beats record
		var upperBounds = [2]int{0, time}
		for {
			holdFor := (upperBounds[0] + upperBounds[1]) / 2
			distance := holdFor * (time - holdFor)
			if distance >= record {
				upperBounds[0] = holdFor
			} else {
				upperBounds[1] = holdFor
			}
			if upperBounds[1]-upperBounds[0] <= 1 {
				break
			}
		}

		sum *= upperBounds[1] - lowerBounds[1]
	}

	return sum
}

func doPartTwo(input []byte) int {

	// Times on the first lije
	var race [2]int
	var tempT int
	var tempD int

	lineLength := len(input) / 2

	var pos int
	for {
		c1 := input[pos]
		c2 := input[pos+lineLength]

		// If c1 is a space or num and c2 is a space or num
		if (c1 == ' ' || c1 >= '0' && c1 <= '9') && (c2 == ' ' || c2 >= '0' && c2 <= '9') && (c1 != ' ' || c2 != ' ') {
			if c1 != ' ' {
				tempT = tempT*10 + int(c1-'0')
			}
			if c2 != ' ' {
				tempD = tempD*10 + int(c2-'0')
			}
		} else if c1 == '\n' && c2 == '\n' {
			race = [2]int{tempT, tempD}
			break
		}
		pos++
	}

	// For each race (time and record distance)
	time := race[0]
	record := race[1]

	// find lowest bound that beats record
	var lowerBounds = [2]int{0, time}
	for {
		holdFor := (lowerBounds[0] + lowerBounds[1]) / 2
		distance := holdFor * (time - holdFor)
		if distance > record {
			lowerBounds[1] = holdFor
		} else {
			lowerBounds[0] = holdFor
		}
		if lowerBounds[1]-lowerBounds[0] <= 1 {
			break
		}
	}

	// find upper bound that beats record
	var upperBounds = [2]int{0, time}
	for {
		holdFor := (upperBounds[0] + upperBounds[1]) / 2
		distance := holdFor * (time - holdFor)
		if distance >= record {
			upperBounds[0] = holdFor
		} else {
			upperBounds[1] = holdFor
		}
		if upperBounds[1]-upperBounds[0] <= 1 {
			break
		}
	}

	times := upperBounds[1] - lowerBounds[1]

	return times
}

func doPartOneMath(input []byte) int {

	// Times on the first lije
	var races [][2]int
	var tempT int
	var tempD int

	lineLength := len(input) / 2

	var pos int
	for {
		c1 := input[pos]
		c2 := input[pos+lineLength]

		// If c1 is a space or num and c2 is a space or num
		if (c1 == ' ' || c1 >= '0' && c1 <= '9') && (c2 == ' ' || c2 >= '0' && c2 <= '9') && (c1 != ' ' || c2 != ' ') {
			if c1 != ' ' {
				tempT = tempT*10 + int(c1-'0')
			}
			if c2 != ' ' {
				tempD = tempD*10 + int(c2-'0')
			}
		} else if c1 == ' ' && c2 == ' ' || c1 == '\n' && c2 == '\n' {
			if tempD != 0 || tempT != 0 {
				races = append(races, [2]int{tempT, tempD})
			}
			tempT = 0
			tempD = 0
			if c1 == '\n' && c2 == '\n' {
				break
			}
		}
		pos++
	}

	var sum int = 1

	// For each race (time and record distance)
	for _, race := range races {
		time := race[0]
		record := race[1]

		b := time
		c := -record
		a := -1

		det := b*b - 4*a*c

		x1 := float64(-float64(b)+math.Sqrt(float64(det))) / float64(2*a)
		x2 := float64(-float64(b)-math.Sqrt(float64(det))) / float64(2*a)

		score := int(math.Floor(x2) - math.Ceil(x1) + 1)

		sum *= score
	}

	return sum
}

func doPartTwoMath(input []byte) int {

	// Times on the first lije
	var race [2]int
	var tempT int
	var tempD int

	lineLength := len(input) / 2

	var pos int
	for {
		c1 := input[pos]
		c2 := input[pos+lineLength]

		// If c1 is a space or num and c2 is a space or num
		if (c1 == ' ' || c1 >= '0' && c1 <= '9') && (c2 == ' ' || c2 >= '0' && c2 <= '9') && (c1 != ' ' || c2 != ' ') {
			if c1 != ' ' {
				tempT = tempT*10 + int(c1-'0')
			}
			if c2 != ' ' {
				tempD = tempD*10 + int(c2-'0')
			}
		} else if c1 == '\n' && c2 == '\n' {
			race = [2]int{tempT, tempD}
			break
		}
		pos++
	}

	// For each race (time and record distance)
	time := race[0]
	record := race[1]

	b := time
	c := -record
	a := -1

	det := b*b - 4*a*c

	x1 := float64(-float64(b)+math.Sqrt(float64(det))) / float64(2*a)
	x2 := float64(-float64(b)-math.Sqrt(float64(det))) / float64(2*a)

	return int(math.Floor(x2) - math.Ceil(x1) + 1)
}
