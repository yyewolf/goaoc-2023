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

func doPartOne(file []byte) int {
	buffer := bufio.NewReader(bytes.NewBuffer(file))

	var sum int

	for l, _, err := buffer.ReadLine(); err == nil; l, _, err = buffer.ReadLine() {
		first := -1
		last := -1
		for _, c := range l {
			I := int(c - '0')
			if I >= 0 && I <= 9 {
				first = I
				break
			}
		}
		for i := len(l) - 1; i >= 0; i-- {
			I := int(l[i] - '0')
			if I >= 0 && I <= 9 {
				last = I
				break
			}
		}
		sum += first<<3 + first<<1 + last
	}

	return sum
}

func fastNumbers(b []byte, N *int) {
	switch b[len(b)-1] {
	case 'o':
		switch b[len(b)-2] {
		case 'w':
			if b[len(b)-3] == 't' {
				*N = 2
			}
		case 'r':
			if b[len(b)-3] == 'e' && b[len(b)-4] == 'z' {
				*N = 0
			}
		}
	case 'e':
		switch b[len(b)-2] {
		case 'n':
			switch b[len(b)-3] {
			case 'i':
				if b[len(b)-4] == 'n' {
					*N = 9
				}
			case 'o':
				*N = 1
			}
		case 'v':
			if b[len(b)-3] == 'i' && b[len(b)-4] == 'f' {
				*N = 5
			}
		case 'e':
			if b[len(b)-3] == 'r' && b[len(b)-4] == 'h' && b[len(b)-5] == 't' {
				*N = 3
			}
		}
	case 'n':
		if b[len(b)-2] == 'e' && b[len(b)-3] == 'v' && b[len(b)-4] == 'e' && b[len(b)-5] == 's' {
			*N = 7
		}
	case 't':
		if b[len(b)-2] == 'h' && b[len(b)-3] == 'g' && b[len(b)-4] == 'i' && b[len(b)-5] == 'e' {
			*N = 8
		}
	case 'x':
		if b[len(b)-2] == 'i' && b[len(b)-3] == 's' {
			*N = 6
		}
	case 'r':
		if b[len(b)-2] == 'u' && b[len(b)-3] == 'o' && b[len(b)-4] == 'f' {
			*N = 4
		}
	}
}

func doPartTwo(file []byte) int {
	buffer := bufio.NewReader(bytes.NewBuffer(file))

	var sum int

	for l, _, err := buffer.ReadLine(); err == nil; l, _, err = buffer.ReadLine() {

		first := -1
		b := make([]byte, 6)
		for i, c := range l {
			if i >= 6 {
				b = l[i-6 : i+1]
			} else {
				b = append(b, c)
			}
			fastNumbers(b, &first)

			if first != -1 {
				break
			}

			I := int(c - '0')
			if I >= 0 && I <= 9 {
				first = I
				break
			}
		}
		last := -1
		for i := len(l) - 1; i >= 0; i-- {
			if i >= 6 {
				b = l[i-6 : i+1]
			} else {
				b = l[:i+1]
			}
			fastNumbers(b, &last)

			if last != -1 {
				break
			}

			I := int(l[i] - '0')
			if I >= 0 && I <= 9 {
				last = I
				break
			}
		}
		sum += first<<3 + first<<1 + last
	}

	return sum
}
