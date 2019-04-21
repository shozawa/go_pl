package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	pre, post := outliner()
	forEachNode(doc, pre, post)
}

func outliner() (pre, post func(*html.Node)) {
	var depth int
	pre = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth++
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
	}
	post = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
			depth--
		}
	}
	return
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if n == nil {
		return
	}
	pre(n)
	if n.FirstChild != nil {
		forEachNode(n.FirstChild, pre, post)
	}
	if n.NextSibling != nil {
		forEachNode(n.NextSibling, pre, post)
	}
	post(n)
}
