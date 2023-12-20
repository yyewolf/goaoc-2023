package cli

import (
	"aocli/template/internal/folder"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cmdRT = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Run the tests for the current day (short: t)",
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
		// Tell the use we're running the input, put "test" in red
		fmt.Println("🎄 Running with \033[31mtest\033[0m input for day", day, "of year", year, ":")

		// Chdir to PWD
		os.Chdir(os.Getenv("PWD"))

		// Executes `go run . test`
		c := exec.Command("go", "run", ".", "test")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		// Prints the output of `go run . test`
		fmt.Print(string(out))
	},
}

var cmdRI = &cobra.Command{
	Use:     "input",
	Aliases: []string{"i"},
	Short:   "Run the input for the current day (short: i)",
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
		// Tell the use we're running the input, put "real" in green
		fmt.Println("🎄 Running with \033[32mreal\033[0m input for day", day, "of year", year, ":")

		// Chdir to PWD
		os.Chdir(os.Getenv("PWD"))

		// Executes `go run . test`
		c := exec.Command("go", "run", ".")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		// Prints the output of `go run . test`
		fmt.Print(string(out))
	},
}

var cmdR = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Run the current day (short: r)",
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
		cmdRT.Run(cmd, args)
	},
}

func init() {
	cmdR.AddCommand(cmdRT)
	cmdR.AddCommand(cmdRI)

	rootCmd.AddCommand(cmdR)
}
