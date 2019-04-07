package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shozawa/go_pl/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		fmt.Println("error")
	}

	issues := (*result).Items

	fmt.Print("within one month\n")
	for _, issue := range github.FilterIssues(issues, isWithinOneMonth) {
		printIssue(issue)
	}

	fmt.Print("within a year\n")
	for _, issue := range github.FilterIssues(issues, isWithinOneYear) {
		printIssue(issue)
	}

	fmt.Print("over a year\n")
	for _, issue := range github.FilterIssues(issues, isOverOneYear) {
		printIssue(issue)
	}
}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}

// TODO: 判定関数群のテストとリファクタリング
func isWithinOneMonth(issue *github.Issue) bool {
	return issue.CreatedAt.After(time.Now().AddDate(0, -1, 0))
}

func isWithinOneYear(issue *github.Issue) bool {
	if isWithinOneMonth(issue) {
		return false
	}

	return issue.CreatedAt.After(time.Now().AddDate(-1, 0, 0))
}

func isOverOneYear(issue *github.Issue) bool {
	if isWithinOneMonth(issue) || isWithinOneYear(issue) {
		return false
	} else {
		return true
	}
}
