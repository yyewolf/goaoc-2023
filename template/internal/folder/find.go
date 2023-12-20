package folder

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetYearFolder() string {
	var pwd = os.Getenv("PWD")

	// pwd looks like ***/year/day01 or ***/year/
	splt := strings.Split(pwd, "/")

	if len(splt) < 2 {
		return ""
	}

	// Check if the last folder is a year
	re := regexp.MustCompile(`\d{4}`)
	if re.MatchString(splt[len(splt)-1]) {
		return splt[len(splt)-1]
	}

	if re.MatchString(splt[len(splt)-2]) {
		return splt[len(splt)-2]
	}

	return ""
}

func GetDayFolder() int {
	var pwd = os.Getenv("PWD")

	// pwd looks like ***/year/day01
	splt := strings.Split(pwd, "/")

	if len(splt) < 2 {
		return 0
	}

	// Check if the last folder is a day
	re := regexp.MustCompile(`day\d{2}`)
	if re.MatchString(splt[len(splt)-1]) {
		// Get the last two characters
		lastTwo := splt[len(splt)-1][len(splt[len(splt)-1])-2:]
		day, _ := strconv.Atoi(lastTwo)
		return day
	}

	return 0
}

func FindRoot() string {
	var at = os.Getenv("PWD")

	for {
		if _, err := os.Stat(at + "/.git"); os.IsNotExist(err) {
			at = strings.TrimSuffix(at, "/"+strings.Split(at, "/")[len(strings.Split(at, "/"))-1])
		} else {
			break
		}
	}

	return at
}
