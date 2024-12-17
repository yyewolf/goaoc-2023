package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func do(A int) int {
	B := A % 8
	B ^= 1
	C := A / 32
	A = A / 8
	B ^= C
	B ^= 6
	return B % 8
}

func calcA(p []int) (out int) {
	for pos, val := range p {
		out += val << (pos * 3)
	}
	return out
}

func doPartTwo(input string) int {
	lines := strings.Split(input, "\n")

	var initialMachine machine
	fmt.Sscanf(lines[0], "Register A: %d", &initialMachine.A)
	fmt.Sscanf(lines[1], "Register B: %d", &initialMachine.B)
	fmt.Sscanf(lines[2], "Register C: %d", &initialMachine.C)

	rawProgram := strings.Split(strings.Split(lines[4], ": ")[1], ",")

	var program []int64
	for _, number := range rawProgram {
		parsed, _ := strconv.ParseInt(number, 10, 64)
		program = append(program, parsed)
	}

	// rawProgramCheck := strings.Join(rawProgram, ",")

	var possibilities = make([]int, len(program))
	for pos := range possibilities {
		possibilities[pos] = 1
	}

	for pos := len(program) - 1; pos >= 4; pos-- {
		for i := 0; i < 8; i++ {
			possibilities[pos] = i

			var m = new(machine)
			*m = initialMachine
			m.A = int64(calcA(possibilities))
			out := strings.Split(m.run(program), ",")
			if pos > len(out)-1 {
				continue
			}

			if out[pos] == rawProgram[pos] {
				break
			}
		}
	}

	var out = []string{"0"}
	A := int64(calcA(possibilities)) + 10000000

	for !slices.Equal(out, rawProgram) {
		A++

		m := new(machine)
		*m = initialMachine
		m.A = A
		out = strings.Split(m.run(program), ",")
	}

	return int(A)
}
