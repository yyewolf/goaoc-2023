package main

import (
	"aocli/template/internal/config"
	"aocli/template/internal/folder"
	"html/template"
	"os"
)

// Parse template from file : .config.secrets.template
var secretsTemplate = template.Must(template.ParseFiles("template/.config.secrets.template"))
var configTemplate = template.Must(template.ParseFiles("template/.config.template"))

func saveSecrets() {
	// Create file
	file, err := os.Create(config.SecretsPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write to file
	err = secretsTemplate.Execute(file, config.C.Secrets)
	if err != nil {
		panic(err)
	}
}

func saveConfig() {
	// Create file
	file, err := os.Create(config.ConfigPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write to file
	err = configTemplate.Execute(file, config.C.Public)
	if err != nil {
		panic(err)
	}

	folder.CreateDay(config.C.Public.CurrentYear, 1)
}
