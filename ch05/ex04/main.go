package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	visit(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {
			if src := getSrc(n); src != "" {
				fmt.Println(src)
			}
		}
	})
}

func getSrc(n *html.Node) string {
	switch n.Data {
	case "a", "link":
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	case "img", "script":
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				return attr.Val
			}
		}
	}
	return ""
}

func visit(n *html.Node, f func(*html.Node)) {
	if n == nil {
		return
	}
	f(n)
	if n.FirstChild != nil {
		visit(n.FirstChild, f)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling, f)
	}
}
