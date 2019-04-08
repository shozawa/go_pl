package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/shozawa/go_pl/ch04/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	report := template.Must(template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
