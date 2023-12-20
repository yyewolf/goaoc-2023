package markdown

import (
	"aocli/template/internal/aoc"
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
)

var (
	formatBadgeLatestCommit = "![](https://img.shields.io/github/last-commit/%s/%s?style=flat-square)"
	formatBadgeCurrentDay   = "![](https://img.shields.io/badge/day 📅-%s-blue)"
	formatBadgeStars        = "![](https://img.shields.io/badge/stars ⭐-%d-yellow)"
	formatBadgeDays         = "![](https://img.shields.io/badge/days completed-%d-red)"
)

func GenerateBadges(year string) (res string) {
	// Get current day
	day := aoc.DefaultDay()

	currentYear := time.Now().Year()

	currentDayBadge := ""

	if day < 25 && year == fmt.Sprintf("%d", currentYear) {
		currentDayBadge = fmt.Sprintf(formatBadgeCurrentDay+"\n", day)
	}

	// Get stars
	stars, err := aoc.GetAllStars(year)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	// Count days with 2 stars
	days := 0
	starCount := 0
	for _, v := range stars {
		if v == 2 {
			days++
		}
		starCount += v
	}

	// Print badges
	currentStarsBadge := fmt.Sprintf(formatBadgeStars, starCount)
	daysBadge := fmt.Sprintf(formatBadgeDays, days)

	// Get current git remote url
	gitRepo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	r, err := gitRepo.Remote("origin")
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	// r can be git@X:username/repo.git
	// or
	// r can be https://X/yyewolf/goaoc.git

	var currentCommitsBadge string

	if strings.HasPrefix(r.Config().URLs[0], "git@") {
		// Remove git@ and : and replace / with -
		important := strings.Split(r.Config().URLs[0], ":")[1]
		username := strings.Split(important, "/")[0]
		repo := strings.Split(important, "/")[1]
		repo = strings.ReplaceAll(repo, ".git", "")
		currentCommitsBadge = fmt.Sprintf(formatBadgeLatestCommit, username, repo)
	} else {
		// Remove https:// and replace / with -
		important := strings.Split(r.Config().URLs[0], "https://")[1]
		username := strings.Split(important, "/")[1]
		repo := strings.Split(important, "/")[2]
		repo = strings.ReplaceAll(repo, ".git", "")
		currentCommitsBadge = fmt.Sprintf(formatBadgeLatestCommit, username, repo)
	}

	res = fmt.Sprintf("%s\n%s%s\n%s", currentCommitsBadge, currentDayBadge, currentStarsBadge, daysBadge)

	// replace spaces with %20
	res = strings.ReplaceAll(res, " ", "%20")

	return
}
