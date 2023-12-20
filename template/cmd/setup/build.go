package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setupAskForBuildAndInstall() {
	// Ask user if they want to build or install the CLI with boxes
	fmt.Printf("🔨 Do you want to build and/or install the CLI?\n")

	var options = []string{"Build and install", "Build", "No"}
	var selected = 0

	// Read byte by byte from os.Stdin without pressing enter
	enableRawMode()
	defer disableRawMode()

	for {
		// Hide cursor
		fmt.Printf("\033[?25l")

		// Print options
		for i, option := range options {
			fmt.Printf("\033[K")
			if i == selected {
				fmt.Printf("🔳 %s\n", option)
			} else {
				fmt.Printf("⬜️ %s\n", option)
			}
		}

		// listen for arrow and enter keys
		key := getKey()

		var exit = false
		switch key {
		case arrowUp:
			if selected > 0 {
				selected--
			} else {
				selected = len(options) - 1
			}
		case arrowDown:
			if selected < len(options)-1 {
				selected++
			} else {
				selected = 0
			}
		case enter:
			// Return selected
			exit = true
		}

		if exit {
			break
		}

		// Go back len(options) lines
		fmt.Printf("\033[%dA", len(options))
	}

	// Show cursor
	fmt.Printf("\033[?25h")

	currentDir := os.Getenv("PWD")
	switch selected {
	case 2:
		return
	case 1:
		// Build the CLI
		fmt.Println("🔨 Building the CLI...")
		cmd := exec.Command("go", "build", "-o", "template/bin/goaoc", fmt.Sprintf("%s/template/cmd/cli/main.go", currentDir))
		cmd.Run()
	case 0:
		// Build and install the CLI
		fmt.Println("🔨 Building and installing the CLI...")
		cmd := exec.Command("go", "build", "-o", "template/bin/goaoc", fmt.Sprintf("%s/template/cmd/cli/main.go", currentDir))
		cmd.Run()

		// print logs

		cmd = exec.Command("sudo", "mv", "template/bin/goaoc", "/usr/local/bin/goaoc")
		cmd.Run()

		// Check if it is in path
		_, err := exec.LookPath("goaoc")
		if err != nil {
			fmt.Println("❌ It looks like goaoc is not in your path. Please add /usr/local/bin to your path.")
		} else {
			fmt.Println("✅ goaoc is now in your path. You can run it from anywhere by typing goaoc.")
		}
	}
}
