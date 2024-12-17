package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type instruction int64

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type machine struct {
	A int64
	B int64
	C int64

	ip int
}

func (m *machine) instruct(opcode instruction, operand int64) *int64 {
	// Do operand
	var comboOperandValue int64 = operand

	switch comboOperandValue {
	case 4:
		comboOperandValue = m.A
	case 5:
		comboOperandValue = m.B
	case 6:
		comboOperandValue = m.C
	}

	skip := false
	defer func() {
		if !skip {
			m.ip += 2
		}
	}()

	switch opcode {
	// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
	case adv:
		m.A = int64(float64(m.A) / math.Pow(2.0, float64(comboOperandValue)))
	// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
	case bxl:
		m.B ^= operand
	// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
	case bst:
		m.B = comboOperandValue % 8
	// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
	case jnz:
		if m.A == 0 {
			break
		}
		m.ip = int(operand)
		skip = true
	// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
	case bxc:
		m.B ^= m.C
	// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
	case out:
		var out = comboOperandValue % 8
		return &out
	// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
	case bdv:
		m.B = int64(float64(m.A) / math.Pow(2.0, float64(comboOperandValue)))
	// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
	case cdv:
		m.C = int64(float64(m.A) / math.Pow(2.0, float64(comboOperandValue)))
	}

	return nil
}

func (m *machine) run(program []int64) string {
	var outputs []string
	for {
		if m.ip < 0 || m.ip > len(program)-2 {
			break
		}

		opcode, operand := program[m.ip], program[m.ip+1]
		output := m.instruct(instruction(opcode), operand)
		if output != nil {
			outputs = append(outputs, fmt.Sprint(*output))
		}
	}

	return strings.Join(outputs, ",")
}

func doPartOne(input string) string {
	lines := strings.Split(input, "\n")

	var machine machine
	fmt.Sscanf(lines[0], "Register A: %d", &machine.A)
	fmt.Sscanf(lines[1], "Register B: %d", &machine.B)
	fmt.Sscanf(lines[2], "Register C: %d", &machine.C)

	rawProgram := strings.Split(strings.Split(lines[4], ": ")[1], ",")

	var program []int64
	for _, number := range rawProgram {
		parsed, _ := strconv.ParseInt(number, 10, 64)
		program = append(program, parsed)
	}

	return machine.run(program)
}
