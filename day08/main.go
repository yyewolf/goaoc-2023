package main

import (
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

var network [0x1000000]int

func doPartOne(input []byte) int {
	// First line is the key, L = 0, R = 1
	var key []int
	var pos int

	for i := range input {
		if input[i] == 'L' {
			key = append(key, 0)
		} else if input[i] == 'R' {
			key = append(key, 1)
		}
		if input[i] == '\n' {
			break
		}
		pos++
	}

	// Second line is empty
	pos += 2

	for ; pos < len(input); pos++ {
		ai := int(input[pos])<<16 + int(input[pos+1])<<8 + int(input[pos+2])
		bi := int(input[pos+7])<<16 + int(input[pos+8])<<8 + int(input[pos+9])
		ci := int(input[pos+12])<<16 + int(input[pos+13])<<8 + int(input[pos+14])

		di := bi + ci<<24

		network[ai] = di
		pos += 16
	}

	var steps int

	// Count how many steps it takes to reach "ZZZ" using the key as a pattern
	at := int('A')<<16 + int('A')<<8 + int('A')
	end := int('Z')<<16 + int('Z')<<8 + int('Z')
	for at != end {
		at = (network[at] >> (key[steps%len(key)] * 24)) & 0xFFFFFF
		steps++
	}

	return steps
}

// gcd calculates the greatest common divisor using the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the least common multiple using the GCD for multiple inputs.
func lcm(numbers ...int) int {
	if len(numbers) < 2 {
		panic("At least two numbers are required to calculate LCM.")
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result = (result * num) / gcd(result, num)
	}
	return result
}

func doPartTwo(input []byte) int {
	// First line is the key, L = 0, R = 1
	var key []int
	var pos int

	for i := range input {
		if input[i] == 'L' {
			key = append(key, 0)
		} else if input[i] == 'R' {
			key = append(key, 1)
		}
		if input[i] == '\n' {
			break
		}
		pos++
	}

	// Second line is empty
	pos += 2

	var at []int

	for ; pos < len(input); pos++ {
		ai := int(input[pos])<<16 + int(input[pos+1])<<8 + int(input[pos+2])
		bi := int(input[pos+7])<<16 + int(input[pos+8])<<8 + int(input[pos+9])
		ci := int(input[pos+12])<<16 + int(input[pos+13])<<8 + int(input[pos+14])

		if ai&0xFF == 0x41 {
			at = append(at, ai)
		}

		network[ai] = bi + ci<<24
		pos += 16
	}

	var steps int

	// Count how many steps it takes to reach "ZZZ" using the key as a pattern
	var cycles []int
	var lenAt = len(at)
	var total int
	for {
		k := key[steps%len(key)]
		for i := 0; i < len(at); i++ {
			total++
			at[i] = (network[at[i]] >> (k * 24)) & 0xFFFFFF
			// Check if it ends with Z
			if at[i]%256 == 0x5A {
				cycles = append(cycles, steps+1)
				// remove from at
				at = append(at[:i], at[i+1:]...)
				i--
			}
		}
		steps++
		if len(cycles) == lenAt {
			break
		}
	}

	return lcm(cycles...)
}
