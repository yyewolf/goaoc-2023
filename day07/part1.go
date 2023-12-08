package main

import (
	"sort"
)

func (card *Card) _type() int {
	// If there's only one letter, it's a five of a kind
	if len(card.CardsNoDupe) == 1 {
		return 1
	}
	// If there's only two letters, it's either a four of a kind or a full house
	if len(card.CardsNoDupe) == 2 {
		// If the first letter is repeated four times, it's a four of a kind
		countF := 0
		countS := 0
		for _, c := range card.Cards {
			if c == card.CardsNoDupe[0] {
				countF++
			} else if c == card.CardsNoDupe[1] {
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
	if len(card.CardsNoDupe) == 3 {
		// If the first letter is repeated three times, it's a three of a kind
		countF := 0
		countS := 0
		countT := 0
		for _, c := range card.Cards {
			if c == card.CardsNoDupe[0] {
				countF++
			} else if c == card.CardsNoDupe[1] {
				countS++
			} else if c == card.CardsNoDupe[2] {
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

	if len(card.CardsNoDupe) == 4 {
		return 6
	}

	return 7
}

func (c *Card) less(t *Card) bool {
	return c.score < t.score
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
	for _, c := range input {
		switch {
		case c == '\n':
			cards[card].score += (8 - cards[card]._type()) << 20
			state = false
			card++
		case c == ' ':
			state = true
		case state && c >= '0' && c <= '9':
			cards[card].Bid = cards[card].Bid*10 + int(c-'0')
		default:
			// Only append if not in Cards
			seen := false
			for _, c2 := range cards[card].Cards {
				if c2 == c {
					seen = true
					break
				}
			}
			if !seen {
				cards[card].CardsNoDupe = append(cards[card].CardsNoDupe, c)
			}
			cards[card].Cards = append(cards[card].Cards, c)
			cards[card].score = cards[card].score<<4 + remap[int(c)]
		}
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].less(&cards[j])
	})

	sum := 0

	for i, c := range cards {
		sum += c.Bid * (i + 1)
	}

	// if sum != 248453531 {
	// 	fmt.Println("Wrong answer:", sum)
	// }

	return sum
}
