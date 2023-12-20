package main

import (
	"aocli/template/internal/cli"
	"aocli/template/internal/folder"
	"aocli/template/internal/markdown"
	"os"
)

var root = folder.FindRoot()

func init() {
	// Change working directory to root
	err := os.Chdir(root)
	if err != nil {
		panic(err)
	}

	markdown.InitTemplates()
}

func main() {
	// folder.CreateDay("2023", "01")
	// fmt.Println(folder.GetYearFolder())
	// fmt.Println(folder.GetBenches("2015"))
	// fmt.Println(markdown.GenerateBenches("2015"))
	// fmt.Println()
	// fmt.Println(markdown.GenerateStars("2015"))
	// fmt.Println(markdown.GenerateReadme("2015"))

	cli.Execute()
}
