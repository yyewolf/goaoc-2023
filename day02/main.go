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

// [3]int{r,g,b}
func fastParser(input *[]byte, output *[3]int) {
	in := *input
	o := *output
	o[0] = 0
	o[1] = 0
	o[2] = 0
	var tempc = -1
	var temp = -1
	var tempPow = 1
	for i := len(in) - 1; i >= 0; i-- {
		// read until comma
		// ignore spaces
		if in[i] == ' ' {
			continue
		}
		if in[i] == ';' || in[i] == ':' {
			// remove from end to here
			in = in[:i]
			*input = in
			break
		}
		if in[i] == ',' {
			o[tempc] = temp
			tempc = -1
			temp = -1
			tempPow = 1
			continue
		}
		if tempc == -1 {
			switch in[i] {
			case 'd':
				tempc = 0
			case 'n':
				tempc = 1
			case 'e':
				tempc = 2
			}
			continue
		}
		if temp == -1 {
			temp = 0
		}
		c := int(in[i] - '0')
		if c < 0 || c > 9 {
			continue
		}
		temp += c * tempPow
		tempPow *= 10
	}
	o[tempc] = temp
	*output = o
}

func doPartOne(input []byte) int {
	sum := 0
	readBuffer := bufio.NewReader(bytes.NewReader(input))
	i := 0
	for {
		line, _, err := readBuffer.ReadLine()
		if err != nil {
			break
		}
		i++

		isPossible := true
		var rgb [3]int
		for {
			fastParser(&line, &rgb)
			if rgb[0] == 0 && rgb[1] == 0 && rgb[2] == 0 {
				break
			}
			if rgb[0] > 12 || rgb[1] > 13 || rgb[2] > 14 {
				isPossible = false
				break
			}
		}

		if isPossible {
			sum += i
		}
	}

	return sum
}

func doPartTwo(input []byte) int {
	sum := 0
	readBuffer := bufio.NewReader(bytes.NewReader(input))
	i := 0
	for {
		line, _, err := readBuffer.ReadLine()
		if err != nil {
			break
		}
		i++
		var max [3]int = [3]int{0, 0, 0}

		var rgb [3]int
		for {
			fastParser(&line, &rgb)
			if rgb[0] == 0 && rgb[1] == 0 && rgb[2] == 0 {
				break
			}
			if rgb[0] > max[0] {
				max[0] = rgb[0]
			}
			if rgb[1] > max[1] {
				max[1] = rgb[1]
			}
			if rgb[2] > max[2] {
				max[2] = rgb[2]
			}
		}

		sum += max[0] * max[1] * max[2]
	}

	return sum
}
