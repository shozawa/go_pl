package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/shozawa/go_pl/ch04/github"
)

const templ = `
	<h1>{{.TotalCount}} issues</h1>
	<table>
	  <tr style='text=align: left'>
	    <th>#</th>
	    <th>State</th>
	    <th>User</th>
	    <th>Title</th>
	  </tr>
	  {{range .Items}}
	  <tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td>{{.User.Login}}</td>
		<td>{{.Title}}</td>
	  </tr>
	  {{end}}
	</table>
`

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
