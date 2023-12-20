package cli

import (
	"aocli/template/internal/markdown"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

var cmdWorkflow = &cobra.Command{
	Use:   "workflow",
	Short: "Run the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		// List folders that match the pattern "20[0-9][0-9]" but are not the current years
		folders, err := os.ReadDir(".")
		if err != nil {
			panic(err)
		}

		re := regexp.MustCompile(`20[0-9][0-9]`)
		var years []int
		for _, folder := range folders {
			if folder.IsDir() && re.MatchString(folder.Name()) {
				year, _ := strconv.Atoi(folder.Name())
				years = append(years, year)
			}
		}

		sort.Ints(years)

		// Generate all readmes and put them in ./year/README.md
		// Last one also goes in ./README.md
		for _, year := range years {
			readme := markdown.GenerateReadme(fmt.Sprintf("%d", year))
			err := os.WriteFile(fmt.Sprintf("%d/README.md", year), []byte(readme), 0644)
			if err != nil {
				panic(err)
			}

			if year == years[len(years)-1] {
				err := os.WriteFile("README.md", []byte(readme), 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdWorkflow)
}
