package main

import (
	"os"
)

// parse reads the file given in input and return the elves
func parse(file string) (string, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
