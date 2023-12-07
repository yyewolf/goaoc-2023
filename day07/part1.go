package main

import "sort"

func (c *Card) _type() int {
	// Remove duplicate letters in cards
	var temp []byte
	for _, c := range c.Cards {
		if len(temp) == 0 {
			temp = append(temp, c)
			continue
		}
		var found = false
		for _, t := range temp {
			if t == c {
				found = true
				break
			}
		}
		if !found {
			temp = append(temp, c)
		}
	}

	// If there's only one letter, it's a five of a kind
	if len(temp) == 1 {
		return 1
	}
	// If there's only two letters, it's either a four of a kind or a full house
	if len(temp) == 2 {
		// If the first letter is repeated four times, it's a four of a kind
		countF := 0
		countS := 0
		for _, c := range c.Cards {
			if c == temp[0] {
				countF++
			} else if c == temp[1] {
				countS++
			}
		}
		if countF == 4 || countS == 4 {
			return 2
		}
		// Otherwise, it's a full house
		return 3
	}

	// If there's only three letters, it's either a three of a kind or a two pair
	if len(temp) == 3 {
		// If the first letter is repeated three times, it's a three of a kind
		countF := 0
		countS := 0
		countT := 0
		for _, c := range c.Cards {
			if c == temp[0] {
				countF++
			} else if c == temp[1] {
				countS++
			} else if c == temp[2] {
				countT++
			}
		}
		m := max(countF, countS, countT)
		if m == 3 {
			return 4
		}
		if m == 2 {
			return 5
		}
		return 5
	}

	if len(temp) == 4 {
		return 6
	}

	return 7
}

var order = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func (c *Card) less(t *Card) bool {
	cType := c._type()
	tType := t._type()
	if cType < tType {
		return false
	}
	if cType > tType {
		return true
	}

	for i := 0; i < len(c.Cards); i++ {
		var cVal, tVal int
		for j, o := range order {
			if c.Cards[i] == o {
				cVal = j
			}
			if t.Cards[i] == o {
				tVal = j
			}
		}
		if cVal == tVal {
			continue
		}
		return cVal < tVal
	}

	return false
}

func doPartOne(input []byte) int {
	// Parse first
	var cards []*Card
	var temp = &Card{}
	var state bool
	for _, c := range input {
		if c == '\n' {
			cards = append(cards, temp)
			temp = &Card{}
			state = false
			continue
		}
		if c == ' ' {
			state = true
			continue
		}
		if state && c >= '0' && c <= '9' {
			temp.Bid = temp.Bid*10 + int(c-'0')
		} else {
			temp.Cards = append(temp.Cards, c)
		}
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].less(cards[j])
	})

	sum := 0

	for i, c := range cards {
		sum += c.Bid * (i + 1)
	}

	return sum
}
