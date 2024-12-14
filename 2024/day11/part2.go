package main

import (
	"strconv"
	"strings"
)

func doPartTwo(input string) int {
	var stones = make(map[[2]int]int)
	fields := strings.Fields(input)
	for _, field := range fields {
		val := fastAtoi(field)
		stones[[2]int{val, 0}]++
	}

	var blinkAmount int = 75
	// key is the stone, val is the amount of stones with that key
	for blink := range blinkAmount {
		for key, amount := range stones {
			stone := key[0]
			stoneBlink := key[1]
			if stoneBlink != blink {
				continue
			}
			delete(stones, key)
			if stone == 0 {
				stones[[2]int{1, blink + 1}] += amount
			} else {
				stoneStr := strconv.Itoa(stone)
				if len(stoneStr)%2 == 0 {
					a := fastAtoi(stoneStr[:len(stoneStr)/2])
					b := fastAtoi(stoneStr[len(stoneStr)/2:])
					stones[[2]int{a, blink + 1}] += amount
					stones[[2]int{b, blink + 1}] += amount
				} else {
					stones[[2]int{stone * 2024, blink + 1}] += amount
				}
			}
		}
	}

	var sum int
	for _, amount := range stones {
		sum += amount
	}

	return sum
}
