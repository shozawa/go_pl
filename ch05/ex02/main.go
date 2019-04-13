package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	count := make(map[string]int)
	count = visit(count, doc)
	for _, elm := range sorted(count) {
		fmt.Printf("%s\t%d\n", elm, count[elm])
	}
}

func visit(count map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return count
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	count = visit(count, n.FirstChild)
	count = visit(count, n.NextSibling)
	return count
}

func sorted(count map[string]int) []string {
	s := make([]string, 0, len(count))
	for k, _ := range count {
		s = append(s, k)
	}
	sort.Slice(s, func(i, j int) bool { return count[s[i]] > count[s[j]] })
	return s
}
