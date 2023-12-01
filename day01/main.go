package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var _ [100000000]byte
}

func main() {
	answer := doPartOne(input)
	fmt.Println(answer)

	answer = doPartTwo(input)
	fmt.Println(answer)
}

func doPartOne(file []byte) int {
	sum := 0

	first := -1
	last := -1
	for _, c := range file {
		I := int(c - '0')
		if I < 0 || I > 9 {
			if c == '\n' {
				sum += first<<3 + first<<1 + last
				first = -1
				last = -1
			}
		} else {
			if first == -1 {
				first = I
				last = I
			} else {
				last = I
			}
		}
	}

	sum += first<<3 + first<<1 + last

	return sum
}

func doFirstLast(first, last *int, I int) {
	if *first == -1 {
		*first = I
		*last = I
	} else {
		*last = I
	}
}

func doPartTwo(file []byte) int {
	sum := 0

	first := -1
	last := -1
	var rotatingBuffer = make([]byte, 6)
	for i, c := range file {
		if i >= 6 {
			rotatingBuffer = file[i-6 : i]
		} else {
			rotatingBuffer = append(rotatingBuffer, c)
		}

		switch rotatingBuffer[len(rotatingBuffer)-1] {
		case 'o':
			switch rotatingBuffer[len(rotatingBuffer)-2] {
			case 'w':
				if rotatingBuffer[len(rotatingBuffer)-3] == 't' {
					doFirstLast(&first, &last, 2)
				}
			case 'r':
				if rotatingBuffer[len(rotatingBuffer)-3] == 'e' && rotatingBuffer[len(rotatingBuffer)-4] == 'z' {
					doFirstLast(&first, &last, 0)
				}
			}
		case 'e':
			switch rotatingBuffer[len(rotatingBuffer)-2] {
			case 'n':
				switch rotatingBuffer[len(rotatingBuffer)-3] {
				case 'i':
					if rotatingBuffer[len(rotatingBuffer)-4] == 'n' {
						doFirstLast(&first, &last, 9)
					}
				case 'o':
					doFirstLast(&first, &last, 1)
				}
			case 'v':
				if rotatingBuffer[len(rotatingBuffer)-3] == 'i' && rotatingBuffer[len(rotatingBuffer)-4] == 'f' {
					doFirstLast(&first, &last, 5)
				}
			case 'e':
				if rotatingBuffer[len(rotatingBuffer)-3] == 'r' && rotatingBuffer[len(rotatingBuffer)-4] == 'h' && rotatingBuffer[len(rotatingBuffer)-5] == 't' {
					doFirstLast(&first, &last, 3)
				}
			}
		case 'n':
			if rotatingBuffer[len(rotatingBuffer)-2] == 'e' && rotatingBuffer[len(rotatingBuffer)-3] == 'v' && rotatingBuffer[len(rotatingBuffer)-4] == 'e' && rotatingBuffer[len(rotatingBuffer)-5] == 's' {
				doFirstLast(&first, &last, 7)
			}
		case 't':
			if rotatingBuffer[len(rotatingBuffer)-2] == 'h' && rotatingBuffer[len(rotatingBuffer)-3] == 'g' && rotatingBuffer[len(rotatingBuffer)-4] == 'i' && rotatingBuffer[len(rotatingBuffer)-5] == 'e' {
				doFirstLast(&first, &last, 8)
			}
		case 'x':
			if rotatingBuffer[len(rotatingBuffer)-2] == 'i' && rotatingBuffer[len(rotatingBuffer)-3] == 's' {
				doFirstLast(&first, &last, 6)
			}
		case 'r':
			if rotatingBuffer[len(rotatingBuffer)-2] == 'u' && rotatingBuffer[len(rotatingBuffer)-3] == 'o' && rotatingBuffer[len(rotatingBuffer)-4] == 'f' {
				doFirstLast(&first, &last, 4)
			}
		}

		I := int(c - '0')
		if I < 0 || I > 9 {
			if c == '\n' {
				sum += first<<3 + first<<1 + last
				first = -1
				last = -1
			}
		} else {
			doFirstLast(&first, &last, I)
		}
	}

	sum += first<<3 + first<<1 + last

	return sum
}
