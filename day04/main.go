package main

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	answer := doPartOne(input)
	fmt.Println(answer)

	answer = doPartTwo(input)
	fmt.Println(answer)
}

var winningNs [100]bool
var empty [100]bool

func fastReset() {
	copy(winningNs[:], empty[:])
}

func doCard(line []byte) int {
	fastReset()
	buf := bufio.NewReader(bytes.NewBuffer(line))

	// read til :
	buf.ReadBytes(':')
	// skip space
	buf.ReadByte()

	// Read winning numbers
	for j := 0; j < 10; j++ {
		b, err := buf.ReadByte()
		if err != nil || b == '|' {
			break
		}
		b2, err := buf.ReadByte()
		if err != nil {
			break
		}
		if b == ' ' {
			winningNs[uint8(b2-'0')] = true
		} else {
			winningNs[uint8(b-'0')*10+uint8(b2-'0')] = true
		}
		buf.ReadByte()
	}

	// skip space
	buf.ReadByte()
	buf.ReadByte()

	var score = 0

	// Read my numbers
	for j := 0; j < 25; j++ {
		b, err := buf.ReadByte()
		if err != nil || b == '|' {
			break
		}
		b2, err := buf.ReadByte()
		if err != nil {
			break
		}
		n := uint8(0)
		if b == ' ' {
			n = uint8(b2 - '0')
		} else {
			n = uint8(b-'0')*10 + uint8(b2-'0')
		}
		if winningNs[n] {
			if score == 0 {
				score = 1
			} else {
				// Multiply by two
				score <<= 1
			}
		}
		buf.ReadByte()
	}
	return score
}

func doPartOne(input []byte) int {
	buf := bufio.NewReader(bytes.NewBuffer(input))

	sum := 0

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		sum += doCard(line)
	}

	return sum
}

func doCardTwo(line []byte) int {
	fastReset()
	buf := bufio.NewReader(bytes.NewBuffer(line))

	// read til :
	buf.ReadBytes(':')
	// skip space
	buf.ReadByte()

	// Read winning numbers
	for j := 0; j < 10; j++ {
		b, err := buf.ReadByte()
		if err != nil || b == '|' {
			break
		}
		b2, err := buf.ReadByte()
		if err != nil {
			break
		}
		if b == ' ' {
			winningNs[uint8(b2-'0')] = true
		} else {
			winningNs[uint8(b-'0')*10+uint8(b2-'0')] = true
		}
		buf.ReadByte()
	}

	// skip space
	buf.ReadByte()
	buf.ReadByte()

	var score = 0

	// Read my numbers
	for j := 0; j < 25; j++ {
		b, err := buf.ReadByte()
		if err != nil || b == '|' {
			break
		}
		b2, err := buf.ReadByte()
		if err != nil {
			break
		}
		n := uint8(0)
		if b == ' ' {
			n = uint8(b2 - '0')
		} else {
			n = uint8(b-'0')*10 + uint8(b2-'0')
		}
		if winningNs[n] {
			score++
		}
		buf.ReadByte()
	}
	return score
}

func doPartTwo(input []byte) int {
	buf := bufio.NewReader(bytes.NewBuffer(input))

	sum := 0

	var i int
	var copies [213]uint32
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		score := doCardTwo(line)

		for j := 0; j < score; j++ {
			copies[i+j+1] += 1 + copies[i]
		}

		sum += score*(1+int(copies[i])) + 1
		i++
	}

	return sum
}
