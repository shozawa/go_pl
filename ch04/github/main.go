package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}

type Issue struct {
	Title string `json:"title"`
}

func SearchIssues() {
	q := url.QueryEscape("repo:golang/go is:open json decoder")
	res, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		// return nil, err
		fmt.Println(err)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	res.Body.Close()
	fmt.Println(*result.Items[0])
}

func main() {
	SearchIssues()
}
