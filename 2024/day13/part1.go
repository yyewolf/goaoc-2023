package main

import (
	"fmt"
	"strings"
)

type game struct {
	Ax, Ay int
	Bx, By int
	Gx, Gy int
}

func (g game) Solve() (Z int) {
	// We want to solve this :
	// (1) A*94 + B*22 = 8400
	// (2) A*94 + B*67 = 5400

	// In this case it would look like :
	// (1)*67 : A*94*67 + B*22*67 = 8400*67
	// (2)*22 : A*94*22 + B*67*22 = 5400*22

	// (1)*67-(2)*22 : A*94*67 - A*94*22 = 8400*67 - 5400*22
	// (1)*67-(2)*22 : A*(94*67 - 94*22) = 8400*67 - 5400*22
	// (1)*67-(2)*22 : A = (8400*67 - 5400*22)/(94*67 - 94*22)

	// Solving from A gives B, we can now generalize this by using maths :

	Δ := g.Ax*g.By - g.Ay*g.Bx
	A := g.By*g.Gx - g.Bx*g.Gy
	B := g.Ay*g.Gx - g.Ax*g.Gy

	// We check if A ≡ Δ && B ≡ Δ
	if A%Δ == 0 && B%Δ == 0 {
		return (A/Δ)*3 + B/(-Δ)
	}

	return 0
}

func doPartOne(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var sum int

	for len(lines) > 0 {
		var game game
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &game.Ax, &game.Ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &game.Bx, &game.By)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &game.Gx, &game.Gy)
		lines = lines[3:]
		if len(lines) != 0 {
			lines = lines[1:]
		}

		sum += game.Solve()
	}

	return sum
}
