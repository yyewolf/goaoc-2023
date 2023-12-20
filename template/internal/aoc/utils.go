package aoc

import (
	"aocli/template/internal/config"
	"aocli/template/internal/folder"
	"fmt"
	"time"
)

func DefaultYear() string {
	var currentYear = time.Now().Year()
	var year = fmt.Sprintf("%d", currentYear)

	if config.C.Public.CurrentYear != "" {
		year = config.C.Public.CurrentYear
	}

	yF := folder.GetYearFolder()
	if yF != "" {
		year = yF
	}

	return year
}

func DefaultDay() int {
	var day = time.Now().Day()

	dF := folder.GetDayFolder()
	if dF != 0 {
		day = dF
	}

	return day
}
