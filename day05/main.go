package main

// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(inputTest)
	println(answer)
}

func doPartOne(input []byte) int {
	var pos = 7
	var temp int
	var j int
	var preEntities []int
	for {
		c := input[pos]
		if c >= '0' && c <= '9' {
			temp = temp*10 + int(c-'0')
		} else {
			preEntities = append(preEntities, temp)
			temp = 0
			if c == '\n' {
				break
			}
		}
		pos++
		j++
	}

	for e := 0; e < 7; e++ {
		pos += 2

		for {
			c := input[pos]
			if c == '\n' {
				break
			}
			pos++
		}

		// seed to soil is destination, source and length of range
		// if seed is > source and < source + length, then it is in range
		// seed becomes seed - source + destination
		var eToE [3]int
		var entities []int
		// Convert to fertilizer
		j = 0
		for {
			if pos >= len(input) {
				break
			}
			c := input[pos]
			if c >= '0' && c <= '9' {
				temp = temp*10 + int(c-'0')
			} else {
				eToE[j] = temp
				temp = 0
				j++
				if c == '\n' {
					for i := 0; i < len(preEntities); i++ {
						seed := preEntities[i]
						if seed >= eToE[1] && seed < eToE[1]+eToE[2] {
							entities = append(entities, seed-eToE[1]+eToE[0])
							preEntities = append(preEntities[:i], preEntities[i+1:]...)
							i--
						}
					}
					if input[pos+1] == '\n' {
						break
					}
					j = 0
				}
			}
			pos++
		}

		// Move all entities to preEntities
		preEntities = append(preEntities, entities...)
		entities = nil
	}

	lowest := 1<<63 - 1
	for _, e := range preEntities {
		if e < lowest {
			lowest = e
		}
	}

	return lowest
}

func checkRangeValid(r [2]int) bool {
	return r[0] < r[1] && r[0] >= 0 && r[1] >= 0
}

func doPartTwo(input []byte) int {
	var pos = 7
	var temp int
	var j int
	var preEntities []int
	for {
		c := input[pos]
		if c >= '0' && c <= '9' {
			temp = temp*10 + int(c-'0')
		} else {
			preEntities = append(preEntities, temp)
			temp = 0
			if c == '\n' {
				break
			}
		}
		pos++
		j++
	}

	// range is start and end
	var preRanges [][2]int
	for i := 0; i < len(preEntities)/2; i++ {
		preRanges = append(preRanges, [2]int{preEntities[i*2], preEntities[i*2] + preEntities[i*2+1]})
	}

	for e := 0; e < 7; e++ {
		pos += 2

		for {
			c := input[pos]
			if c == '\n' {
				break
			}
			pos++
		}

		var eToE [3]int
		var ranges [][2]int

		var doit = func() {
			for i := 0; i < len(preRanges); i++ {
				a := preRanges[i][0]
				b := preRanges[i][1]
				c := eToE[1]
				d := eToE[1] + eToE[2]
				e := eToE[0]

				if a < c && b < c || a > d && b > d {
					continue
				}

				if a > c && a < d && b > c && b < d {
					during := [2]int{e + (a - c), e + ((a - c) + (b - a))}
					if checkRangeValid(during) {
						ranges = append(ranges, during)
					}
					preRanges = append(preRanges[:i], preRanges[i+1:]...)
					i--
				} else if a > c && a < d && b >= d {
					during := [2]int{e + (a - c), e + ((a - c) + (d - a))}
					if checkRangeValid(during) {
						ranges = append(ranges, during)
					}

					after := [2]int{d + 1, b}
					if checkRangeValid(after) {
						preRanges = append(preRanges, after)
					}
					preRanges = append(preRanges[:i], preRanges[i+1:]...)
					i--
				} else if a <= c && b > c && b < d {
					before := [2]int{a, c - 1}
					if checkRangeValid(before) {
						preRanges = append(preRanges, before)
					}

					during := [2]int{e, e + (b - c)}
					if checkRangeValid(during) {
						ranges = append(ranges, during)
					}
					preRanges = append(preRanges[:i], preRanges[i+1:]...)
					i--
				} else if a <= c && b >= d {
					before := [2]int{a, c - 1}
					if checkRangeValid(before) {
						preRanges = append(preRanges, before)
					}

					during := [2]int{e, e + (d - c)}
					if checkRangeValid(during) {
						ranges = append(ranges, during)
					}

					after := [2]int{d + 1, b}
					if checkRangeValid(after) {
						preRanges = append(preRanges, after)
					}
					preRanges = append(preRanges[:i], preRanges[i+1:]...)
					i--
				}
			}
		}

		j = 0
		for {
			if pos >= len(input) {
				eToE[j] = temp
				doit()
				break
			}
			c := input[pos]
			if c >= '0' && c <= '9' {
				temp = temp*10 + int(c-'0')
			} else {
				eToE[j] = temp
				temp = 0
				j++
				if c == '\n' {
					if eToE[0] != 0 || eToE[1] != 0 || eToE[2] != 0 {
						doit()
					}
					if input[pos+1] == '\n' {
						break
					}
					j = 0
				}
			}
			pos++
		}

		// Move all entities to preEntities

		preRanges = append(preRanges, ranges...)
		ranges = nil
	}

	lowest := 1<<63 - 1
	for _, e := range preRanges {
		if e[0] < lowest {
			lowest = e[0]
		}
	}

	return lowest
}
