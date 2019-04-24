package main

import (
	"fmt"
	"log"
	neturl "net/url"
	"os"

	"github.com/shozawa/go_pl/ch05/graph"
	"github.com/shozawa/go_pl/ch05/links"
)

var urls []string = os.Args[1:]

func main() {
	graph.BFS(crawl, urls)
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Exract(url, isSameHost(urls, url))
	if err != nil {
		log.Print(err)
	}
	return list
}

// TODO: コレクション操作を外出しする
func isSameHost(urls []string, url string) bool {
	a, err := neturl.Parse(url)
	if err != nil {
		return false
	}
	for _, e := range urls {
		b, err := neturl.Parse(e)
		if err != nil {
			continue
		}
		if a.Hostname() == b.Hostname() {
			return true
		}
	}
	return false
}
