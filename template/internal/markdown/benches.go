package markdown

import (
	"aocli/template/internal/benches"
	"aocli/template/internal/folder"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

var benchesMd string
var benchesColumnCount int

func initBenches() {
	d, err := os.ReadFile(benchesTemplate)
	if err != nil {
		panic("🚨 benches.md template not found")
	}

	benchesMd = string(d)

	// Search for "BenchesColumns=" in benches.md
	// If not found, panic
	i := strings.Index(benchesMd, "BenchesColumns=")
	if i == -1 {
		panic("🚨 BenchesColumns not found in benches.md")
	}

	// Get column count
	benchesColumnCount, _ = strconv.Atoi(benchesMd[i+15 : i+16])

	// Fill benchesMd with what's in between {{START}} and {{END}}
	// If not found, panic
	start := strings.Index(benchesMd, "{{START}}")
	end := strings.Index(benchesMd, "{{END}}")

	if start == -1 || end == -1 {
		panic("🚨 {{START}} or {{END}} not found in benches.md")
	}

	benchesMd = benchesMd[start+10 : end-1]
}

func GenerateBenches(year string) string {
	var lines = make([]string, int(math.Ceil(25.0/float64(benchesColumnCount))))

	var bs = folder.GetBenches(year)

	// Fill lines
	for i, b := range bs {
		line := i % int(math.Ceil(25.0/float64(benchesColumnCount)))

		lines[line] += fmt.Sprintf("| [Day %d](/%s/day%02d) | ", i+1, year, i+1)

		// Add part one bench if exists
		if len(b.Lines) == 0 {
			lines[line] += "| "
			continue
		}

		p1 := b.Lines[0]
		timeFormatFunc := b.TimeFormatFunc(p1.NsPerOp)
		time := timeFormatFunc(p1.NsPerOp)
		lines[line] += fmt.Sprintf("`%s`", time)

		if (b.Measured & parse.MBPerS) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatMegaBytesPerSecond(p1))
		}
		if (b.Measured & parse.AllocedBytesPerOp) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatBytesAllocPerOp(p1))
		}
		if (b.Measured & parse.AllocsPerOp) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatAllocsPerOp(p1))
		}

		if len(b.Lines) == 1 {
			continue
		}

		lines[line] += " | "

		p2 := b.Lines[1]
		timeFormatFunc = b.TimeFormatFunc(p2.NsPerOp)
		time = timeFormatFunc(p2.NsPerOp)
		lines[line] += fmt.Sprintf("`%s`", time)

		if (b.Measured & parse.MBPerS) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatMegaBytesPerSecond(p2))
		}
		if (b.Measured & parse.AllocedBytesPerOp) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatBytesAllocPerOp(p2))
		}
		if (b.Measured & parse.AllocsPerOp) > 0 {
			lines[line] += fmt.Sprintf("/`%s`", benches.FormatAllocsPerOp(p2))
		}
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

	res = strings.ReplaceAll(res, "/op", "")

	return benchesMd + "\n" + res + "\n"
}
