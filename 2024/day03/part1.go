package main

import (
	"bufio"
	"fmt"
	"strings"
)

// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

// Try to scanf the string for mul(%d,%d)

func doPartOne(input string) int {
	buffer := bufio.NewReader(strings.NewReader(input))

	var result int
	for {
		rune, _, err := buffer.ReadRune()
		if err != nil {
			break
		}
		if rune == 'm' {
			var a, b int
			_, err := fmt.Fscanf(buffer, "ul(%d,%d)", &a, &b)
			if err != nil {
				continue
			}

			// Verify that a and b are within 1-999
			if a < 1 || a > 999 || b < 1 || b > 999 {
				continue
			}
			result += a * b
		}
	}

	return result
}
