package main

import (
	"math/big"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input, nil)
	println(answer)
}

// 26425 is 0b110011100111001 which is 'ZZZ' (the biggest possible)
var network [26426]uint32

func doPartOne(input []byte) int {
	var pos int
	// The key is stored in only 5 uint64s, so we can use a bit pattern to represent it
	var key = make([]uint64, 0, 5)

	var tempInt uint64
	var tempIntPos int
	for i := range input {
		if input[i] == '\n' {
			break
		}

		tempInt |= uint64(input[i]>>4) & 1 << tempIntPos
		tempIntPos++
		if tempIntPos == 64 {
			key = append(key, tempInt)
			tempInt = 0
			tempIntPos = 0
		}
		pos++
	}
	key = append(key, tempInt)

	// Second line is empty
	pos += 2

	// I use the network and
	for ; pos < len(input); pos++ {
		ai := uint32(input[pos]-0x41)<<10 + uint32(input[pos+1]-0x41)<<5 + uint32(input[pos+2]-0x41)
		bi := uint32(input[pos+7]-0x41)<<10 + uint32(input[pos+8]-0x41)<<5 + uint32(input[pos+9]-0x41)
		ci := uint32(input[pos+12]-0x41)<<25 + uint32(input[pos+13]-0x41)<<20 + uint32(input[pos+14]-0x41)<<15

		network[ai] = bi + ci
		pos += 16
	}

	var steps int

	// Count how many steps it takes to reach "ZZZ" using the key as a pattern
	at := uint32(0)
	end := uint32(0b110011100111001)

	var keyPos int
	var keyPosInByte int

	for at != end {
		at = network[at]
		if key[keyPos]&(1<<keyPosInByte) == 0 {
			at &= 32767
		} else {
			at >>= 15
		}
		steps++

		keyPosInByte++
		if keyPosInByte == 64 || (keyPos == len(key)-1 && keyPosInByte == tempIntPos) {
			keyPosInByte = 0
			keyPos++
		}
		if keyPos == len(key) {
			keyPos = 0
		}
	}

	return steps
}

// gcd calculates the Greatest Common Divisor using Euclid's algorithm.
func gcd(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

// lcm calculates the Least Common Multiple using the formula lcm(a, b) = |a * b| / gcd(a, b).
func lcm(a, b *big.Int) *big.Int {
	if a.Sign() == 0 || b.Sign() == 0 {
		return big.NewInt(0)
	}
	g := gcd(a, b)
	return new(big.Int).Abs(new(big.Int).Div(new(big.Int).Mul(a, b), g))
}

// lcmArray calculates the LCM of an array of integers.
func lcmArray(arr []int) *big.Int {
	result := big.NewInt(1)
	for _, num := range arr {
		result = lcm(result, big.NewInt(int64(num)))
	}
	return result
}

func doPartTwo(input []byte, b *testing.B) int {
	var pos int
	var key = make([]uint64, 0, 5)

	var tempInt uint64
	var tempIntPos int
	for i := range input {
		if input[i] == '\n' {
			break
		}

		tempInt |= uint64(input[i]>>4) & 1 << tempIntPos
		tempIntPos++
		if tempIntPos == 64 {
			key = append(key, tempInt)
			tempInt = 0
			tempIntPos = 0
		}
		pos++
	}
	key = append(key, tempInt)

	// Second line is empty
	pos += 2

	var at = make([]uint32, 0, 10)

	for ; pos < len(input); pos++ {
		ai := uint32(input[pos]-0x41)<<10 + uint32(input[pos+1]-0x41)<<5 + uint32(input[pos+2]-0x41)
		bi := uint32(input[pos+7]-0x41)<<10 + uint32(input[pos+8]-0x41)<<5 + uint32(input[pos+9]-0x41)
		ci := uint32(input[pos+12]-0x41)<<26 + uint32(input[pos+13]-0x41)<<21 + uint32(input[pos+14]-0x41)<<16

		if input[pos+2] == 'A' {
			at = append(at, ai)
		}

		network[ai] = bi + ci
		pos += 16
	}

	var steps int

	// Count how many steps it takes to reach "ZZZ" using the key as a pattern
	var cycles = make([]int, 0, 10)
	var lenAt = len(at)
	var keyPos int
	var keyPosInByte int

	var nM int
	var nS int

	var k bool
	for len(cycles) < lenAt {
		k = key[keyPos]&(1<<keyPosInByte) == 0

		for i := 0; i < len(at); i++ {
			at[i] = network[at[i]]
			if k {
				at[i] &= 32767
				nM++
			} else {
				at[i] >>= 16
				nS++
			}
			// Check if it ends with Z
			if at[i]&31 == 25 {
				cycles = append(cycles, steps+1)
				// remove from at
				at = append(at[:i], at[i+1:]...)
				i--
			}
		}
		steps++

		keyPosInByte++
		if keyPosInByte == 64 || (keyPos == len(key)-1 && keyPosInByte == tempIntPos) {
			keyPosInByte = 0
			keyPos++
		}
		if keyPos == len(key) {
			keyPos = 0
		}
	}

	return int(lcmArray(cycles).Int64())
}

func doParsing(input []byte) uint32 {
	var pos int
	// The key is stored in only 5 uint64s, so we can use a bit pattern to represent it
	var key = make([]uint64, 0, 5)

	var tempInt uint64
	var tempIntPos int
	for i := range input {
		if input[i] == '\n' {
			break
		}

		tempInt |= uint64(input[i]>>4) & 1 << tempIntPos
		tempIntPos++
		if tempIntPos == 64 {
			key = append(key, tempInt)
			tempInt = 0
			tempIntPos = 0
		}
		pos++
	}
	key = append(key, tempInt)

	// Second line is empty
	pos += 2

	var at = make([]uint32, 0, 10)
	// I use the network and
	for ; pos < len(input); pos++ {
		ai := uint32(input[pos]-0x41)<<10 + uint32(input[pos+1]-0x41)<<5 + uint32(input[pos+2]-0x41)
		bi := uint32(input[pos+7]-0x41)<<10 + uint32(input[pos+8]-0x41)<<5 + uint32(input[pos+9]-0x41)
		ci := uint32(input[pos+12]-0x41)<<25 + uint32(input[pos+13]-0x41)<<20 + uint32(input[pos+14]-0x41)<<15

		if input[pos+2] == 'A' {
			at = append(at, ai)
		}

		network[ai] = bi + ci
		pos += 16
	}

	return at[len(at)-1]%256 + uint32(key[len(key)-1])
}
