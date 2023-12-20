package main

import (
	"sort"
)

// func _type2() int {
// 	// Remove duplicate letters in cards
// 	var temp = c.CardsNoDupe

// 	// If there's only one letter, it's a five of a kind
// 	if len(temp) == 1 {
// 		return 1
// 	}
// 	// If there's only two letters, it's either a four of a kind or a full house
// 	if len(temp) == 2 {
// 		// If the first letter is repeated four times, it's a four of a kind
// 		countF := 0
// 		countS := 0
// 		for _, c := range c.CardsNoJokers {
// 			if c == temp[0] {
// 				countF++
// 			} else if c == temp[1] {
// 				countS++
// 			}
// 		}
// 		if countF == 4 || countS == 4 {
// 			return 2
// 		}
// 		// Otherwise, it's a full house
// 		return 3
// 	}

// 	// If there's only three letters, it's either a three of a kind or a two pair
// 	if len(temp) == 3 {
// 		// If the first letter is repeated three times, it's a three of a kind
// 		countF := 0
// 		countS := 0
// 		countT := 0
// 		for _, c := range c.CardsNoJokers {
// 			if c == temp[0] {
// 				countF++
// 			} else if c == temp[1] {
// 				countS++
// 			} else if c == temp[2] {
// 				countT++
// 			}
// 		}
// 		m := max(countF, countS, countT)
// 		if m == 3 {
// 			return 4
// 		}
// 		if m == 2 {
// 			return 5
// 		}
// 		return 5
// 	}

// 	if len(temp) == 4 {
// 		return 6
// 	}

// 	return 7
// }

// var order2 = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

var remap2 = [100]int{
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

var cards2 = make([]Card, 1000)

func doPartTwo(input []byte) int {
	// Parse first
	var state bool
	var card int

	var tempCards []byte
	var tempCardsJ []byte
	var tempCardsNoDupe []byte
	var hasJ bool
	for _, c := range input {
		switch {
		case c == '\n':
			var minType = _type(tempCardsNoDupe, tempCards)

			if hasJ {
				if string(tempCards) == "JJJJJ" {
					minType = 1
				} else {
					// Let's say we have AAKKJ
					// tempCards = AAKKJ
					// tempCardsNoDupe = AK
					// maxOccurence = 2
					// maxOccurenceVal = 12

					var maxOccurence int
					var maxOccurenceVal int = -1
					var maxOccurenceChar byte
					for _, c := range tempCardsNoDupe {
						val := remap2[int(c)]
						count := 0
						for _, c2 := range tempCards {
							if c2 == c {
								count++
							}
						}
						// Better count first, then better val
						if count > maxOccurence || (count == maxOccurence && val > maxOccurenceVal) {
							maxOccurence = count
							maxOccurenceVal = val
							maxOccurenceChar = c
						}
					}

					tempCardsJ = tempCards[:]
					for i, c := range tempCards {
						if c == 'J' {
							tempCardsJ[i] = maxOccurenceChar
						}
					}

					minType = _type(tempCardsNoDupe, tempCardsJ)
				}
			}

			cards2[card].score += (8 - minType) << 20

			tempCards = tempCards[:0]
			tempCardsNoDupe = tempCardsNoDupe[:0]
			state = false
			hasJ = false
			card++
		case c == ' ':
			state = true
		case state && c >= '0' && c <= '9':
			cards2[card].Bid = cards2[card].Bid*10 + int(c-'0')
		default:
			// // Only append if not in Cards
			seen := false
			for _, c2 := range tempCards {
				if c2 == c {
					seen = true
					break
				}
			}
			if !seen && c != 'J' {
				tempCardsNoDupe = append(tempCardsNoDupe, c)
			}
			if c == 'J' {
				hasJ = true
			}
			tempCards = append(tempCards, c)
			cards2[card].score = cards2[card].score<<4 + remap2[int(c)]
		}
	}

	// Sort cards
	sort.Sort(CardSlice(cards2))

	sum := 0

	for i, c := range cards2 {
		sum += c.Bid * (i + 1)
	}

	return sum
}
