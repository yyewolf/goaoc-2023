package cli

import (
	"aocli/template/internal/aoc"
	"aocli/template/internal/folder"
	"fmt"

	"github.com/spf13/cobra"
)

var cmdP = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"p"},
	Short:   "Prepare for the current day (short: p)",
	Run: func(cmd *cobra.Command, args []string) {
		in, err := aoc.GetInput(year, fmt.Sprintf("%d", day))
		if err != nil {
			fmt.Println("🚨 Could not get input")
		}

		folder.CreateDay(year, day)
		folder.CreateDayInput(year, day, in)
	},
}

func init() {
	defaultYear := aoc.DefaultYear()
	defaultDay := aoc.DefaultDay()

	cmdP.Flags().StringVarP(&year, "year", "y", defaultYear, "Specify the year for Advent of Code")
	cmdP.Flags().IntVarP(&day, "day", "d", defaultDay, "Specify the day for Advent of Code")

	rootCmd.AddCommand(cmdP)
}
