package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout

func main() {
	pretty()
}

func pretty() {
	doc, err := html.Parse(in)

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}

	var depth int

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild != nil {
				fmt.Fprintf(out, "%*s<%s>\n", depth*2, "", n.Data)
				depth++
			} else {
				var attr string
				for _, a := range n.Attr {
					attr = attr + fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
				}
				fmt.Fprintf(out, "%*s<%s%s/>\n", depth*2, "", n.Data, attr)
			}
		}

		if n.Type == html.TextNode {
			if text := strings.TrimSpace(n.Data); text != "" {
				fmt.Fprintf(out, "%*s%s\n", depth*2, "", text)
			}
		}
	}

	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild != nil {
				depth--
				fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
			}
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
