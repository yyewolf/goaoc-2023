package main

import (
	"bytes"
	"strconv"
)

func main() {
	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func picksTheorem(points [][]int64) float64 {
	n := len(points)
	if n < 3 {
		// Not a valid polygon
		return 0
	}

	// Calculate the area of the polygon using Pick's theorem
	A := 0.0
	B := int64(0)

	for i := 0; i < n-1; i++ {
		xi, yi := points[i][0], points[i][1]
		xi1, yi1 := points[i+1][0], points[i+1][1]

		A += float64(xi*yi1 - xi1*yi)
		B += gcd(abs(xi-xi1), abs(yi-yi1))
	}

	// Add the contribution from the last and first vertices
	A += float64(points[n-1][0]*points[0][1] - points[0][0]*points[n-1][1])
	B += gcd(abs(points[n-1][0]-points[0][0]), abs(points[n-1][1]-points[0][1]))

	// Apply Pick's theorem
	I := (A + float64(B) + 2) / 2

	return I
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func doPartOne(input []byte) int {
	var x, y int64

	lines := bytes.Split(input, []byte{'\n'})

	var points [][]int64
	for _, line := range lines {
		splt := bytes.Split(line, []byte{' '})
		dir := splt[0]
		dist, _ := strconv.ParseInt(string(splt[1]), 10, 64)

		switch dir[0] {
		case 'R':
			points = append(points, []int64{x + dist, y})
			x += dist
		case 'L':
			points = append(points, []int64{x - dist, y})
			x -= dist
		case 'U':
			points = append(points, []int64{x, y - dist})
			y -= dist
		case 'D':
			points = append(points, []int64{x, y + dist})
			y += dist
		}
	}

	area := picksTheorem(points)

	return int(area)
}

func doPartTwo(input []byte) int {
	var x, y int64

	lines := bytes.Split(input, []byte{'\n'})

	var points [][]int64
	for _, line := range lines {
		splt := bytes.Split(line, []byte{' '})
		dir := splt[2][len(splt[2])-2]
		dist, _ := strconv.ParseInt(string(splt[2][2:len(splt[2])-2]), 16, 64)

		switch dir {
		case '0':
			points = append(points, []int64{x + dist, y})
			x += dist
		case '2':
			points = append(points, []int64{x - dist, y})
			x -= dist
		case '3':
			points = append(points, []int64{x, y - dist})
			y -= dist
		case '1':
			points = append(points, []int64{x, y + dist})
			y += dist
		}
	}

	area := picksTheorem(points)

	return int(area)
}
