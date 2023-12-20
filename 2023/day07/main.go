package main

import (
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Card struct {
	Cards         []byte
	CardsNoDupe   []byte
	CardsNoJokers []byte
	Bid           int
	score         int
}

// CardSlice is a slice of Card structs
type CardSlice []Card

// Len returns the length of the CardSlice
func (cs CardSlice) Len() int {
	return len(cs)
}

// Swap swaps the elements with indexes i and j in the CardSlice
func (cs CardSlice) Swap(i, j int) {
	// Only swap bid and score because we don't care about the cards at this point
	cs[i].Bid, cs[j].Bid = cs[j].Bid, cs[i].Bid
	cs[i].score, cs[j].score = cs[j].score, cs[i].score
}

// Less returns true if the element with index i should sort before the element with index j
func (cs CardSlice) Less(i, j int) bool {
	return cs[i].score < cs[j].score
}
