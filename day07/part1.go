package main

import (
	"sort"
)

func _type(nodupe, cards []byte) int {
	// If there's only one letter, it's a five of a kind
	if len(nodupe) == 1 {
		return 1
	}
	// If there's only two letters, it's either a four of a kind or a full house
	if len(nodupe) == 2 {
		// If the first letter is repeated four times, it's a four of a kind
		countF := 0
		countS := 0
		for _, c := range cards {
			if c == nodupe[0] {
				countF++
			} else if c == nodupe[1] {
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
	if len(nodupe) == 3 {
		// If the first letter is repeated three times, it's a three of a kind
		countF := 0
		countS := 0
		countT := 0
		for _, c := range cards {
			if c == nodupe[0] {
				countF++
			} else if c == nodupe[1] {
				countS++
			} else if c == nodupe[2] {
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

	if len(nodupe) == 4 {
		return 6
	}

	return 7
}

var cards = make([]Card, 1000)
var remap = [100]int{
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
	84: 10,
	74: 11,
	81: 12,
	75: 13,
	65: 14,
}

func doPartOne(input []byte) int {
	// Parse first
	// var cards[card] = Card{}
	var state bool
	var card int

	var tempCards []byte
	var tempCardsNoDupe []byte
	for _, c := range input {
		switch {
		case c == '\n':
			cards[card].score += (8 - _type(tempCardsNoDupe, tempCards)) << 20
			// Empty lists
			tempCards = tempCards[:0]
			tempCardsNoDupe = tempCardsNoDupe[:0]
			state = false
			card++
		case c == ' ':
			state = true
		case state && c >= '0' && c <= '9':
			cards[card].Bid = cards[card].Bid*10 + int(c-'0')
		default:
			// Only append if not in Cards
			seen := false
			for _, c2 := range tempCards {
				if c2 == c {
					seen = true
					break
				}
			}
			if !seen {
				tempCardsNoDupe = append(tempCardsNoDupe, c)
			}
			tempCards = append(tempCards, c)
			cards[card].score = cards[card].score<<4 + remap[int(c)]
		}
	}

	sort.Sort(CardSlice(cards))

	sum := 0

	for i, c := range cards {
		sum += c.Bid * (i + 1)
	}

	return sum
}
