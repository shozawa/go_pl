package main

import (
	"html/template"
	"log"
	"net/http"
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	report := template.Must(template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	q := []string{"repo:golang/go", "is:open"}
	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
