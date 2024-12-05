package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Rules []Rule

func (r Rules) Verify(list []int) bool {
	for _, rule := range r {
		if !rule.Verify(list) {
			return false
		}
	}
	return true
}

func (r Rule) Sort(list []int) []int {
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
		return list
	}
	if first > second {
		list[first], list[second] = list[second], list[first]
	}
	return list
}

func doPartTwo(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var rules Rules
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

	var incorrectPrints [][]int

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
			incorrectPrints = append(incorrectPrints, list)
		}
	}

	// Sort the incorrectPrints using rules
	for idx, list := range incorrectPrints {
		for !rules.Verify(list) {
			for _, rule := range rules {
				incorrectPrints[idx] = rule.Sort(list)
			}
		}
	}

	// Result is the sum of the middle number of the new list
	var result int

	for _, i := range incorrectPrints {
		middle := len(i) / 2
		result += i[middle]
	}

	return result
}
