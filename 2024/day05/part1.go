package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Rule struct {
	First  int
	Second int
}

func (r Rule) Verify(list []int) bool {
	var first, second int
	var seenFirst, seenSecond bool
	for idx, i := range list {
		if i == r.First {
			seenFirst = true
			first = idx
		}
		if i == r.Second {
			seenSecond = true
			second = idx
		}
	}
	if !seenFirst || !seenSecond {
		return true
	}
	if first > second {
		return false
	}
	return true
}

func doPartOne(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var rules []Rule
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		// rules are written as "%d|%d"
		var rule Rule
		_, _ = fmt.Sscanf(line, "%d|%d", &rule.First, &rule.Second)
		rules = append(rules, rule)
	}

	var result int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		// split the line by ","
		var list []int
		for _, s := range strings.Split(line, ",") {
			var i int
			_, _ = fmt.Sscanf(s, "%d", &i)
			list = append(list, i)
		}

		// check if the list is valid
		var valid = true
		for _, rule := range rules {
			if !rule.Verify(list) {
				valid = false
			}
		}
		if !valid {
			continue
		}

		// if the list is valid, add the middle element to the result
		middle := len(list) / 2
		result += list[middle]
	}

	return result
}
