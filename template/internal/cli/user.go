package cli

import (
	"aocli/template/internal/aoc"
	"fmt"

	"github.com/spf13/cobra"
)

var cmdU = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Get your user ID (short: u)",
	Run: func(cmd *cobra.Command, args []string) {
		userid, err := aoc.GetUserID()
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
		}

		if auto {
			fmt.Println(userid)
			return
		}

		fmt.Println("🎄 Your user ID is:", userid)
	},
}

func init() {
	cmdU.Flags().BoolVarP(&auto, "auto", "a", false, "Output for use in scripts")

	rootCmd.AddCommand(cmdU)
}
