package main

import (
    "fmt"
    "os"

	"golang.org/x/net/html"
	"github.com/shozawa/go_pl/ch07/ex04/strings"
)

func main() {
	reader := strings.NewReader(`
	<h1>sample</h1>
	<a href="/link1">link</a>
	<ul>
		<li><a href="/link2">link in linst</a></li>
	</ul>
	`)
	doc, err := html.Parse(reader)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }
    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}