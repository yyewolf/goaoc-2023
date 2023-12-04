package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	answer := doPartOne(input)
	fmt.Println(answer)

	answer = doPartTwo(input)
	fmt.Println(answer)
}

func doCard(line []byte) int {
	var winningNs [10]int
	buf := bufio.NewReader(bytes.NewBuffer(line))

	// read til :
	buf.ReadBytes(':')
	// skip space
	buf.ReadByte()

	// Read winning numbers
	var i int
	for {
		// read til |
		b, err := buf.ReadByte()
		if err != nil {
			break
		}

		if b == '|' {
			break
		}

		// Read winning numbers
		if b >= '0' && b <= '9' {
			winningNs[i] = int(b - '0')
			for {
				b, err := buf.ReadByte()
				if err != nil {
					break
				}
				if b == ' ' {
					i++
					break
				}
				winningNs[i] = winningNs[i]*10 + int(b-'0')
			}
		}
	}

	// skip space
	buf.ReadByte()

	var score = 0

	var temp = 0
	// Read my numbers
	for {
		// read til |
		b, err := buf.ReadByte()
		if err != nil {
			break
		}

		if b == '\n' {
			break
		}

		// Read winning numbers
		if b >= '0' && b <= '9' {
			temp = int(b - '0')
			for {
				b, err := buf.ReadByte()
				if b == ' ' || err != nil {
					for _, n := range winningNs {
						if n == temp {
							if score == 0 {
								score = 1
							} else {
								// Multiply by two
								score <<= 1
							}
							temp = 0
							break
						}
					}
					break
				}
				temp = temp*10 + int(b-'0')
			}
		}
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
	var winningNs [10]int
	buf := bufio.NewReader(bytes.NewBuffer(line))

	// read til :
	buf.ReadBytes(':')
	// skip space
	buf.ReadByte()

	// Read winning numbers
	var i int
	for {
		// read til |
		b, err := buf.ReadByte()
		if err != nil {
			break
		}

		if b == '|' {
			break
		}

		// Read winning numbers
		if b >= '0' && b <= '9' {
			winningNs[i] = int(b - '0')
			for {
				b, err := buf.ReadByte()
				if err != nil {
					break
				}
				if b == ' ' {
					i++
					break
				}
				winningNs[i] = winningNs[i]*10 + int(b-'0')
			}
		}
	}

	// skip space
	buf.ReadByte()

	var score = 0

	var temp = 0
	// Read my numbers
	for {
		// read til |
		b, err := buf.ReadByte()
		if err != nil {
			break
		}

		if b == '\n' {
			break
		}

		// Read winning numbers
		if b >= '0' && b <= '9' {
			temp = int(b - '0')
			for {
				b, err := buf.ReadByte()
				if b == ' ' || err != nil {
					for _, n := range winningNs {
						if n == temp {
							score++
							temp = 0
							break
						}
					}
					break
				}
				temp = temp*10 + int(b-'0')
			}
		}
	}

	return score
}

func doPartTwo(input []byte) int {
	buf := bufio.NewReader(bytes.NewBuffer(input))

	sum := 0

	var i int
	var copies [213]int
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		score := doCardTwo(line)

		for j := 0; j < score; j++ {
			copies[i+j+1] += 1 + copies[i]
		}

		sum += score*(1+copies[i]) + 1
		i++
	}

	return sum
}
