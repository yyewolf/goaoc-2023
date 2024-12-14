package main

import (
	"strconv"
	"strings"
)

// At each blink :

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

func fastAtoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func doPartOne(input string) int {
	var stones []int
	fields := strings.Fields(input)
	for _, field := range fields {
		stones = append(stones, fastAtoi(field))
	}

	var blinkAmount int = 25
	for range blinkAmount {
		var newStones []int
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				stoneStr := strconv.Itoa(stone)
				if len(stoneStr)%2 == 0 {
					a := fastAtoi(stoneStr[:len(stoneStr)/2])
					b := fastAtoi(stoneStr[len(stoneStr)/2:])
					newStones = append(newStones, a, b)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}
		stones = newStones
	}

	return len(stones)
}
