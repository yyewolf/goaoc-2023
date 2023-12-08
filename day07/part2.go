package main

import (
	"sort"
)

func (c *Card) _type2() int {
	// Remove duplicate letters in cards
	var temp = c.CardsNoDupe

	// If there's only one letter, it's a five of a kind
	if len(temp) == 1 {
		return 1
	}
	// If there's only two letters, it's either a four of a kind or a full house
	if len(temp) == 2 {
		// If the first letter is repeated four times, it's a four of a kind
		countF := 0
		countS := 0
		for _, c := range c.CardsNoJokers {
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
		for _, c := range c.CardsNoJokers {
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

// var order2 = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

var remap2 = [100]int{
	0:  74,
	1:  50,
	2:  51,
	3:  52,
	4:  53,
	5:  54,
	6:  55,
	7:  56,
	8:  57,
	9:  84,
	10: 81,
	11: 75,
	12: 65,

	74: 0,
	50: 1,
	51: 2,
	52: 3,
	53: 4,
	54: 5,
	55: 6,
	56: 7,
	57: 8,
	84: 9,
	81: 10,
	75: 11,
	65: 12,
}

func doPartTwo(input []byte) int {
	// Parse first
	var temp = Card{}
	var state bool
	var card int
	for _, c := range input {
		switch {
		case c == '\n':
			// Add CardsNoJokers
			temp.CardsNoJokers = make([]byte, len(temp.Cards))
			copy(temp.CardsNoJokers, temp.Cards)

			// Check if hasJ
			hasJ := false
			for _, c := range temp.Cards {
				if c == 'J' {
					hasJ = true
					break
				}
			}
			if hasJ {
				var occurences = make(map[byte]int)
				for _, c := range temp.Cards {
					if c == 'J' {
						continue
					}
					occurences[c]++
				}
				// Keep highest occurence
				var maxOccurence = 0
				for _, o := range occurences {
					if o > maxOccurence {
						maxOccurence = o
					}
				}
				// Delete all occurences below maxOccurence
				for k, o := range occurences {
					if o < maxOccurence {
						delete(occurences, k)
					}
				}
				// Pick best card to replace Joker with
				var best byte
				for i := 12; i >= 0; i-- {
					o := remap2[i]
					_, ok := occurences[byte(o)]
					if ok {
						best = byte(o)
						break
					}
				}
				for i, c := range temp.Cards {
					if c == 'J' {
						temp.CardsNoJokers[i] = best
					}
				}

				if len(temp.CardsNoDupe) == 0 {
					temp.CardsNoDupe = append(temp.CardsNoDupe, best)
				}
			}

			temp.score += (8 - temp._type2()) << 20
			cards[card] = temp

			state = false
			temp.Bid = 0
			temp.Cards = []byte{}
			temp.CardsNoDupe = []byte{}
			temp.CardsNoJokers = []byte{}
			temp.score = 0
			card++
		case c == ' ':
			state = true
		case state && c >= '0' && c <= '9':
			temp.Bid = temp.Bid*10 + int(c-'0')
		default:
			// // Only append if not in Cards
			seen := false
			for _, c2 := range temp.Cards {
				if c2 == c {
					seen = true
					break
				}
			}
			if !seen && c != 'J' {
				temp.CardsNoDupe = append(temp.CardsNoDupe, c)
			}
			temp.Cards = append(temp.Cards, c)
			temp.score = temp.score<<4 + remap2[int(c)]
		}
	}

	// Sort cards
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].less(&cards[j])
	})

	sum := 0

	for i, c := range cards {
		sum += c.Bid * (i + 1)
	}

	return sum
}
