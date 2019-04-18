package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for _, t := range visit(nil, doc) {
		fmt.Println(t)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil || n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	// FXIME: 空のテキストノードは取り除いたほうがいいと思う
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	texts = visit(texts, n.FirstChild)
	texts = visit(texts, n.NextSibling)
	return texts
}
