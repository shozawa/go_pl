package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		// err
	}
	e := elementById(doc, "topbar")
	fmt.Println(e)
}

func elementById(doc *html.Node, id string) *html.Node {
	var element *html.Node
	visit(doc, func(n *html.Node) bool {
		if _id := getId(n); _id == id {
			element = n
		}
		return element != nil
	}, nil)
	return element
}

func getId(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "id" {
			return attr.Val
		}
	}
	return ""
}

func visit(n *html.Node, pre, post func(n *html.Node) bool) {
	if n == nil {
		return
	}
	if pre != nil && pre(n) {
		return
	}
	if n.FirstChild != nil {
		visit(n.FirstChild, pre, post)
	}
	if n.NextSibling != nil {
		visit(n.NextSibling, pre, post)
	}
	// [question]
	// post は本処理実行後に呼ばれるので継続するかどうかを見なくていいのでは？
	if post != nil {
		post(n)
	}
}
