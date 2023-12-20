package main

import (
	"bytes"
)

func main() {
	answer := doPartOne(inputTest)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func doPartOne(input []byte) int {
	sum := 0
	var currentHash uint8
	for _, c := range input {
		if c == ',' {
			sum += int(currentHash * 17)
			currentHash = 0
		} else {
			currentHash = currentHash*17 + c
		}
	}

	return sum
}

type hval struct {
	label []byte
	focal uint8
}

var defaultHval = hval{[]byte{}, 0}
var hashmap [256][]hval

func doPartTwo(input []byte) int {
	clear(hashmap[:])

	var label uint8
	var state bool
	var isEquals bool
	var start int
	for i, c := range input {
		if c == ',' {
			label = label * 17
			labelBytes := input[start : i-1]
			if isEquals {
				labelBytes = input[start : i-2]
			}
			focal := input[i-1] - '0'

			if isEquals {
				var f bool
				for i, h := range hashmap[label] {
					if bytes.Equal(h.label, labelBytes) {
						// replace
						// h.focal = focal
						hashmap[label][i].focal = focal

						f = true
					}
				}

				if !f {
					defaultHval.label = labelBytes
					defaultHval.focal = focal
					hashmap[label] = append(hashmap[label], defaultHval)
				}
			} else {
				// Find the label
				for i, h := range hashmap[label] {
					if bytes.Equal(h.label, labelBytes) {
						// remove
						hashmap[label] = append(hashmap[label][:i], hashmap[label][i+1:]...)
					}
				}
			}

			label = 0
			state = false
			isEquals = false
			start = i + 1
		} else {
			if c == '=' || c == '-' {
				state = true
				isEquals = c == '='
			} else if !state {
				label = label*17 + c
			}
		}
	}

	var sum int
	for i, h := range hashmap {
		for j, v := range h {
			sum += (i + 1) * (j + 1) * int(v.focal)
		}
	}

	return sum
}
