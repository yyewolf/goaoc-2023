package cli

import (
	"aocli/template/internal/benches"
	"aocli/template/internal/config"
	"aocli/template/internal/folder"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var cmdB = &cobra.Command{
	Use:     "bench",
	Aliases: []string{"b"},
	Short:   "Run the bench for the current day (short: b)",
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
		// Chdir to PWD
		os.Chdir(os.Getenv("PWD"))

		// Get config for bench flags
		rawFlags := config.C.Public.BenchFlags
		flags := strings.Split(rawFlags, " ")

		command := []string{"go", "test", "-bench", "."}
		command = append(command, flags...)

		fmt.Printf("🎄 Benching day %d (%s):\n", day, year)

		// Executes `go run . test`
		c := exec.Command(command[0], command[1:]...)

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		b, err := benches.ParseGoBenchmark(string(out))
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		fmt.Println(b.String())

		// Save b to file ".bench"
		f, err := os.Create(".bench")
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}
		defer f.Close()

		// Write b as json to file
		d, err := json.Marshal(b)
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}

		_, err = f.Write(d)
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdB)
}
