package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	nodes := ElementByTagName(doc, "h1", "h2", "h3")
	for _, node := range nodes {
		fmt.Println(node.Data)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	visit(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && contains(name, n.Data) {
			nodes = append(nodes, n)
		}
	})
	return nodes
}

func visit(n *html.Node, f func(*html.Node)) {
	if n == nil {
		return
	}
	if f != nil {
		f(n)
		visit(n.FirstChild, f)
		visit(n.NextSibling, f)
	}
}

func contains(list []string, s string) bool {
	for _, elm := range list {
		if elm == s {
			return true
		}
	}
	return false
}
