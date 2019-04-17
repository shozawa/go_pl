package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("words: %d\timages: %d\n", words, images)
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	var text string
	visit(doc, func(n *html.Node) {
		if n.Type == html.TextNode {
			text += n.Data
		}
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
	})
	words = countWords(text)
	return
}

func countWords(text string) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return
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
