package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}

	var depth int

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s/>\n", depth*2, " ", n.Data)
			depth++
		}
	}

	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
		}
	}
	forEachNode(doc, pre, post)
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
