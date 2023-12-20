package main

import (
	"aocli/template/internal/config"
	"fmt"
	"os"
	"strings"
)

func setupAskForSecrets() {
	// Check if .config.secrets exists
	_, err := os.Stat(config.SecretsPath)

	if err == nil {
		// file exists
		fmt.Print("🔑 It looks like secrets are already present. Would you like to update them? (y/N) ")
		var reset string
		fmt.Scanln(&reset)

		reset = strings.ToLower(reset)
		if reset != "y" {
			return
		}

		config.C.Secrets.AocSession = os.Getenv("AOC_SESSION")
	}

	// Ask user for session cookie
	if config.C.Secrets.AocSession == "" {
		fmt.Printf("🍪 Please enter your session cookie: ")
	} else {
		fmt.Printf("🍪 Please enter your session cookie (%s): ", config.C.Secrets.AocSession)
	}
	var session string
	fmt.Scanln(&session)

	if session != "" {
		config.C.Secrets.AocSession = session
	}
}
