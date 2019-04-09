package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	http.HandleFunc("/favicon.ico", faviconHandler)
	// TODO: マイルストーン、ユーザーの一覧を表示する
	r.HandleFunc("/{user}/{repository}", handler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// FIXME: favicon表示されてない？
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: このファイルからの相対パスで指定する
	http.ServeFile(w, r, "./ch04/issureport/favicon.ico")
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	repository := vars["repository"]
	q := []string{fmt.Sprintf("repo:%s/%s", user, repository), "is:open"}

	report := template.Must(template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	result, err := github.SearchIssues(q)

	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
