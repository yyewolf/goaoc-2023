package markdown

import (
	"aocli/template/internal/aoc"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var starsMd string
var starsColumnCount int

func initStars() {
	d, err := os.ReadFile(starsTemplate)
	if err != nil {
		panic("🚨 stars.md template not found")
	}

	starsMd = string(d)

	// Search for "StarsColumns=" in stars.md
	// If not found, panic
	i := strings.Index(starsMd, "StarsColumns=")
	if i == -1 {
		panic("🚨 StarsColumns not found in stars.md")
	}

	// Get column count
	starsColumnCount, _ = strconv.Atoi(starsMd[i+13 : i+14])

	// Fill starsMd with what's in between {{START}} and {{END}}
	// If not found, panic
	start := strings.Index(starsMd, "{{START}}")
	end := strings.Index(starsMd, "{{END}}")

	if start == -1 || end == -1 {
		panic("🚨 {{START}} or {{END}} not found in stars.md")
	}

	starsMd = starsMd[start+10 : end-1]
}

func GenerateStars(year string) string {
	// Get stars
	stars, err := aoc.GetAllStars(year)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	// Fill stars if len < 25
	if len(stars) < 25 {
		for i := len(stars); i < 25; i++ {
			stars = append(stars, 0)
		}
	}

	var lines = make([]string, int(math.Ceil(25.0/float64(starsColumnCount))))

	// Fill lines
	for i, s := range stars {
		line := i % int(math.Ceil(25.0/float64(starsColumnCount)))

		var emoji string
		switch s {
		case 0:
			// emoji waiting
			emoji = "🎄"
		case 1:
			emoji = "🌟"
		case 2:
			emoji = "🌟🌟"
		}
		lines[line] += fmt.Sprintf("| [Day %d](https://adventofcode.com/%s/day/%d) | %s ", i+1, year, i+1, emoji)
	}

	// Fill empty lines
	for i, l := range lines {
		if l == "" {
			for j := 0; j < benchesColumnCount; j++ {
				lines[i] += "| | "
			}
		}
	}

	// Finish lines
	for i, l := range lines {
		lines[i] = l + "|"
	}

	// Join lines
	res := strings.Join(lines, "\n")

	return starsMd + "\n" + res + "\n"
}

func GenerateCompactStars(year string) string {
	// Get stars
	stars, err := aoc.GetAllStars(year)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	// Fill stars if len < 25
	if len(stars) < 25 {
		for i := len(stars); i < 25; i++ {
			stars = append(stars, 0)
		}
	}

	var res = fmt.Sprintf("| [Advent Of Code %s](/%s) | ", year, year)

	// Fill lines
	for _, s := range stars {
		var emoji string
		switch s {
		case 0:
			// emoji waiting
			emoji = "🖤"
		case 1:
			emoji = "💙"
		case 2:
			emoji = "💛"
		}

		res += emoji + " "
	}

	res += "|\n"

	return res
}
