package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Part one
	lines, err := parse("input.txt")
	if err != nil {
		panic(err)
	}

	answer := doPartOne(lines)
	fmt.Println(answer)

	answer = doPartTwo(lines)
	fmt.Println(answer)
}

func doPartOne(file string) int {
	r := regexp.MustCompile(`[^\d\n]+`)
	file = r.ReplaceAllString(file, "")
	lines := strings.Split(file, "\n")
	sum := 0
	for _, line := range lines {
		first := line[0]
		last := line[len(line)-1]
		I, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		sum += I
	}
	return sum
}

var NumToWord = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func doPartTwo(file string) int {
	// First we replace spelled numbers with digits
	for k, v := range NumToWord {
		file = strings.ReplaceAll(file, v, fmt.Sprintf("%s%d%s", v, k, v))
	}

	r := regexp.MustCompile(`[^\d\n]+`)
	file = r.ReplaceAllString(file, "")
	lines := strings.Split(file, "\n")
	sum := 0
	for _, line := range lines {
		I, _ := strconv.Atoi(fmt.Sprintf("%c%c", line[0], line[len(line)-1]))
		sum += I
	}
	return sum
}
