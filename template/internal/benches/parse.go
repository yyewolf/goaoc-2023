package benches

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/tools/benchmark/parse"
)

func ParseGoBenchmark(input string) (*BenchOutputGroup, error) {
	lines := strings.Split(input, "\n")

	// Remove the first 5 lines
	lines = lines[4:]

	// Remove the last 2 lines
	lines = lines[:len(lines)-3]

	currentBenchmark := &BenchOutputGroup{}
	for _, line := range lines {
		b, err := ParseLine(line)
		if err != nil {
			return nil, err
		}
		currentBenchmark.AddLine(b)
	}

	return currentBenchmark, nil
}

type BenchOutputGroup struct {
	Lines []*parse.Benchmark
	// Columns which are in use
	Measured int
}

type Table struct {
	MaxLengths []int
	Cells      [][]string
}

func (g *BenchOutputGroup) String() string {
	if len(g.Lines) == 0 {
		return ""
	}
	columnNames := []string{"benchmark", "iter", "time/iter"}
	if (g.Measured & parse.MBPerS) > 0 {
		columnNames = append(columnNames, "throughput")
	}
	if (g.Measured & parse.AllocedBytesPerOp) > 0 {
		columnNames = append(columnNames, "bytes alloc")
	}
	if (g.Measured & parse.AllocsPerOp) > 0 {
		columnNames = append(columnNames, "allocs")
	}
	table := &Table{Cells: [][]string{columnNames}}

	var underlines []string
	for _, name := range columnNames {
		underlines = append(underlines, strings.Repeat("-", len(name)))
	}
	table.Cells = append(table.Cells, underlines)

	for _, line := range g.Lines {
		// Rename for parts
		if strings.Contains(line.Name, "PartOne") {
			line.Name = "Part 1"
		} else if strings.Contains(line.Name, "PartTwo") {
			line.Name = "Part 2"
		}

		timeFormatFunc := g.TimeFormatFunc(line.NsPerOp)
		row := []string{line.Name, FormatIterations(line.N), timeFormatFunc(line.NsPerOp)}
		if (g.Measured & parse.MBPerS) > 0 {
			row = append(row, FormatMegaBytesPerSecond(line))
		}
		if (g.Measured & parse.AllocedBytesPerOp) > 0 {
			row = append(row, FormatBytesAllocPerOp(line))
		}
		if (g.Measured & parse.AllocsPerOp) > 0 {
			row = append(row, FormatAllocsPerOp(line))
		}
		table.Cells = append(table.Cells, row)
	}
	for i := range columnNames {
		maxLength := 0
		for _, row := range table.Cells {
			if len(row[i]) > maxLength {
				maxLength = len(row[i])
			}
		}
		table.MaxLengths = append(table.MaxLengths, maxLength)
	}
	var buf bytes.Buffer
	for _, row := range table.Cells {
		for i, cell := range row {
			var format string
			switch i {
			case 0:
				format = "%%-%ds   "
			case len(row) - 1:
				format = "%%%ds"
			default:
				format = "%%%ds   "
			}
			fmt.Fprintf(&buf, fmt.Sprintf(format, table.MaxLengths[i]), cell)
		}
		fmt.Fprint(&buf, "\n")
	}
	return buf.String()
}

func FormatIterations(iter int) string {
	return strconv.FormatInt(int64(iter), 10)
}

func (g *BenchOutputGroup) TimeFormatFunc(c float64) func(float64) string {
	switch {
	case c < float64(10000*time.Nanosecond):
		return func(ns float64) string {
			return fmt.Sprintf("%.2f ns/op", ns)
		}
	case c < float64(time.Millisecond):
		return func(ns float64) string {
			return fmt.Sprintf("%.2f μs/op", ns/1000)
		}
	case c < float64(10*time.Second):
		return func(ns float64) string {
			return fmt.Sprintf("%.2f ms/op", (ns / 1e6))
		}
	default:
		return func(ns float64) string {
			return fmt.Sprintf("%.2f s/op", ns/1e9)
		}
	}
}

func FormatMegaBytesPerSecond(l *parse.Benchmark) string {
	if (l.Measured & parse.MBPerS) == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f MB/s", l.MBPerS)
}

func FormatBytesAllocPerOp(l *parse.Benchmark) string {
	if (l.Measured & parse.AllocedBytesPerOp) == 0 {
		return ""
	}
	return fmt.Sprintf("%d B/op", l.AllocedBytesPerOp)
}

func FormatAllocsPerOp(l *parse.Benchmark) string {
	if (l.Measured & parse.AllocsPerOp) == 0 {
		return ""
	}
	return fmt.Sprintf("%d allocs/op", l.AllocsPerOp)
}

func (g *BenchOutputGroup) AddLine(line *parse.Benchmark) {
	g.Lines = append(g.Lines, line)
	g.Measured |= line.Measured
}

func ParseLine(line string) (*parse.Benchmark, error) {
	return parse.ParseLine(line)
}
