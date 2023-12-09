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

func fastReset() {
	clear(winningNs)
}

func doCard(line []byte) int {
	fastReset()

	var pos int
	// read til :
	for i := 0; i < len(line); i++ {
		if line[i] == ':' {
			pos = i + 2
			break
		}
	}

	// Read winning numbers
	for j := 0; j < 10; j++ {
		b1 := line[pos+j]
		if b1 == '|' {
			break
		}
		b2 := line[pos+j+1]
		if b1 == ' ' {
			winningNs[uint8(b2-'0')] = true
		} else {
			winningNs[uint8(b1-'0')*10+uint8(b2-'0')] = true
		}
		pos += 2
	}

	pos += 2

	var score = 0

	// Read my numbers
	for j := 10; j < 35; j++ {
		if pos+j > len(line)-1 {
			break
		}
		b1 := line[pos+j]
		b2 := line[pos+j+1]
		n := uint8(0)
		if b1 == ' ' {
			n = uint8(b2 - '0')
		} else {
			n = uint8(b1-'0')*10 + uint8(b2-'0')
		}
		if winningNs[n] {
			if score == 0 {
				score = 1
			} else {
				// Multiply by two
				score <<= 1
			}
		}
		pos += 2
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

	var pos int
	// read til :
	for i := 0; i < len(line); i++ {
		if line[i] == ':' {
			pos = i + 2
			break
		}
	}

	// Read winning numbers
	for j := 0; j < 10; j++ {
		b1 := line[pos+j]
		if b1 == '|' {
			break
		}
		b2 := line[pos+j+1]
		if b1 == ' ' {
			winningNs[uint8(b2-'0')] = true
		} else {
			winningNs[uint8(b1-'0')*10+uint8(b2-'0')] = true
		}
		pos += 2
	}

	pos += 2

	var score = 0

	// Read my numbers
	for j := 10; j < 35; j++ {
		if pos+j > len(line)-1 {
			break
		}
		b1 := line[pos+j]
		b2 := line[pos+j+1]
		n := uint8(0)
		if b1 == ' ' {
			n = uint8(b2 - '0')
		} else {
			n = uint8(b1-'0')*10 + uint8(b2-'0')
		}
		if winningNs[n] {
			score++
		}
		pos += 2
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
