package main

import (
	"bufio"
	"bytes"
	"strings"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)	
}

var memo = make(map[int]int)

func countArrangements(row string, groups []int) int {
	clear(memo)

	var backtrack func(int, int) int
	backtrack = func(p, n int) (r int) {
		if v, ok := memo[p<<10+n]; ok {
			return v
		}
		defer func() {
			memo[p<<10+n] = r
		}()

		if p >= len(row) {
			// return n == len(groups)
			if n == len(groups) {
				r = 1
				return
			}
			return
		}

		r = 0

		if row[p] == '.' || row[p] == '?' {
			r = backtrack(p+1, n)
		}

		if row[p] == '#' || row[p] == '?' {
			// verify bounds
			if n >= len(groups) || p+groups[n] >= len(row) {
				return r
			}

			if !strings.Contains(row[p:p+groups[n]], ".") && row[p+groups[n]] != '#' {
				r += backtrack(p+groups[n]+1, n+1)
			}
		}

		return r
	}

	return backtrack(0, 0)
}

func doPartOne(input []byte) int {
	var buffer = bufio.NewReader(bytes.NewReader(input))

	var arrangements int
	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil {
			break
		}

		var row []byte
		var temp int
		var groups []int
		for i, c := range line {
			if c == ' ' {
				row = line[:i]
			} else if len(row) > 0 {
				if c != ',' && c != '\n' {
					temp = temp*10 + int(c-'0')
				} else {
					groups = append(groups, temp)
					temp = 0
				}
			}
		}
		row = append(row, '.')
		count := countArrangements(string(row), groups)
		arrangements += count
	}

	return arrangements
}

func doPartTwo(input []byte) int {
	var buffer = bufio.NewReader(bytes.NewReader(input))

	var arrangements int
	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil {
			break
		}

		var row []byte
		var temp int
		var groups []int
		for i, c := range line {
			if c == ' ' {
				row = line[:i]
			} else if len(row) > 0 {
				if c != ',' && c != '\n' {
					temp = temp*10 + int(c-'0')
				} else {
					groups = append(groups, temp)
					temp = 0
				}
			}
		}

		// Row is now 5 copies of itself, separated by ?
		var first = row[:]
		var firstG = groups[:]
		for i := 0; i < 4; i++ {
			row = append(row, '?')
			row = append(row, first...)
			groups = append(groups, firstG...)
		}
		row = append(row, '.')

		count := countArrangements(string(row), groups)
		arrangements += count
	}

	return arrangements
}
