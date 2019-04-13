package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// FIXME: 題意に沿ってる？
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for _, t := range visit(nil, doc) {
		// TODO: なんか空行がたくさん出るので削除する
		fmt.Println(t)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil || n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	texts = visit(texts, n.FirstChild)
	texts = visit(texts, n.NextSibling)
	return texts
}
