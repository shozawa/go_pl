package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login string
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	var result IssueSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &result, nil
}

func FilterIssues(issues []*Issue, f func(*Issue) bool) []*Issue {
	var o []*Issue
	for _, i := range issues {
		if f(i) {
			o = append(o, i)
		}
	}
	return o
}
