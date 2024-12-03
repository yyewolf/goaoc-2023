package main

import (
	"bufio"
	"fmt"
	"strings"
)

func doPartTwo(input string) int {
	buffer := bufio.NewReader(strings.NewReader(input))

	var result int
	var enabled bool = true
	for {
		rune, _, err := buffer.ReadRune()
		if err != nil {
			break
		}
		if rune == 'd' {
			// Check wheter it says "do()" or "don't()"
			var do string

			// Read the next three runes
			for i := 0; i < 3; i++ {
				rune, _, err := buffer.ReadRune()
				if err != nil {
					break
				}
				do += string(rune)
			}

			if do == "o()" {
				enabled = true
			} else if do == "on'" {
				// check parentheses
				for i := 0; i < 3; i++ {
					rune, _, err := buffer.ReadRune()
					if err != nil {
						break
					}
					do += string(rune)
				}

				if do == "on't()" {
					enabled = false
				}
			}
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
			if !enabled {
				continue
			}
			result += a * b
		}
	}

	return result
}
