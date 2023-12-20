package main

import (
	"aocli/template/internal/aoc"
	"aocli/template/internal/config"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func enableRawMode() {
	var t unix.Termios
	termios.Tcgetattr(uintptr(syscall.Stdin), &t)
	t.Lflag &^= unix.ECHO | unix.ICANON
	termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSANOW, &t)
}

func disableRawMode() {
	var t unix.Termios
	termios.Tcgetattr(uintptr(syscall.Stdin), &t)
	t.Lflag |= unix.ECHO | unix.ICANON
	termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSANOW, &t)
}

func main() {
	// Clear the terminal and print ascii art of go
	fmt.Print("\033[H\033[2J")

	// Do a lil cute animation
	var frame string
	var r, g, b int
	for i := 0; i < len(logo); i++ {
		frame = logo[:i]
		// fmt.Printf("\033[2;0H%s", frame)

		// do random colors :
		if i%100 == 0 {
			r = rand.Intn(255)
			g = rand.Intn(255)
			b = rand.Intn(255)
		}
		// iconic blue for last colors
		if i > len(logo)-100 {
			r = 0
			g = 200
			b = 200
		}

		fmt.Printf("\033[2;0H\033[38;2;%d;%d;%dm%s", r, g, b, frame)

		time.Sleep(1 * time.Microsecond)
	}
	// reset color
	fmt.Printf("\033[2;0H%s", logo)
	fmt.Printf("\033[0m")

	// Skip 2-3 lines
	fmt.Printf("\n\n\n")

	// Check if it has been run before
	setupCheckIfRunBefore()

	// Ask for session cookie
	setupAskForYear()

	// Ask for session cookie
	setupAskForSecrets()

	// Ask if they want to build and install the CLI
	setupAskForBuildAndInstall()

	// Ask if they want to install the github workflow
	setupAskForGithubWorkflow()

	fmt.Println("🔧 Saving config...")
	// Save config
	saveSecrets()
	saveConfig()
}

func setupCheckIfRunBefore() {
	// Check if .config exists
	_, err := os.Stat(".config")

	if err == nil {
		// file exists
		fmt.Print("🔧 It looks like you already ran setup. Would you like to continue? (Y/n) ")
		var reset string
		fmt.Scanln(&reset)

		reset = strings.ToLower(reset)
		if reset == "n" {
			os.Exit(0)
		}
	}
}

func setupAskForYear() {
	var year = aoc.DefaultYear()

	// Ask user for year with emojis
	fmt.Printf("📅 Please enter the year you want to setup for (%s) : ", year)
	fmt.Scanln(&year)

	// Check that it's a valid year (2015-now)
	// If it's not, ask again
	// If it is, continue
	re := regexp.MustCompile(`^20[1-9][0-9]$`)
	if !re.MatchString(year) {
		fmt.Println("❌ Please enter a valid year.")
		setupAskForYear()
	} else {
		config.C.Public.CurrentYear = year
	}
}

func setupAskForGithubWorkflow() {

	// Ask user for year with emojis
	fmt.Printf("📅 Would you like to install the github workflow? (Y/n) : ")
	var install string
	fmt.Scanln(&install)

	// Check that it's a valid year (2015-now)
	// If it's not, ask again
	// If it is, continue
	install = strings.ToLower(install)
	if install == "n" {
		return
	}

	// Copy template/.github/workflows/*.yml to .github/workflows/*.yml
	fmt.Println("🔧 Copying github workflow...")
	err := os.MkdirAll(".github/workflows", 0755)
	if err != nil {
		panic(err)
	}

	// Run cmd
	cmd := exec.Command("cp", "-r", "./template/.github", "./")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// Tell user to put what's in template/.config.secrets and template/.config in Github Secrets at :
	// https://github.com/username/repo/settings/secrets/actions
	fmt.Println("🔧 Please put what's in template/.config.secrets and template/.config in Github Secrets at :")
	fmt.Println("🔧 https://github.com/username/repo/settings/secrets/actions")
}
