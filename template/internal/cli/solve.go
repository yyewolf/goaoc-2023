package cli

import (
	"aocli/template/internal/aoc"
	"aocli/template/internal/folder"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func doSubmit(part int, answer string) {
	fmt.Printf("🎄 Are you sure you want to submit the answer for day %d of %s? (y/N) ", day, year)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input != "y" {
		fmt.Println("🎄 Aborting...")
		return
	}

	fmt.Println("🎄 Submitting...")
	success, err := aoc.Submit(year, day, part, answer)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
		return
	}

	if success {
		fmt.Println("🎄 GG!")
	} else {
		fmt.Println("🚨 This is not the correct answer!")
	}
}

var cmdSolve = &cobra.Command{
	Use:     "solve",
	Aliases: []string{"s"},
	Short:   "Run the input for the current day and submit (short: s)",
	PreRun: func(cmd *cobra.Command, args []string) {
		d := folder.GetDayFolder()
		if d == 0 {
			fmt.Println("🚨 You are not in a day folder!")
			os.Exit(1)
		}

		y := folder.GetYearFolder()
		if y == "" {
			fmt.Println("🚨 You are not in a year folder!")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Ask the user if they're sure
		fmt.Println("🎄 Make sure that these are correct :")

		// Chdir to PWD
		os.Chdir(os.Getenv("PWD"))

		// Executes `go run . test`
		c := exec.Command("go", "run", ".")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		lines := strings.Split(string(out), "\n")

		part1 := strings.TrimSpace(lines[0])
		part2 := strings.TrimSpace(lines[1])

		if part2 != "0" {
			fmt.Println("🎄 Submitting part 2:", part2)
			doSubmit(2, part2)
			return
		}
		if part1 != "0" {
			fmt.Println("🎄 Submitting part 1:", part1)
			doSubmit(1, part1)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdSolve)
}
