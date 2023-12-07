package main

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Card struct {
	Cards         []byte
	CardsNoJokers []byte
	Bid           int
}
