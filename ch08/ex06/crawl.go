package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shozawa/go_pl/ch05/links"
)

var depth = flag.Int("depth", 1, "どのくらいの深さまでリンクをたどるか")

func main() {
	flag.Parse()

	type item struct {
		links []string
		depth int
	}
	worklist := make(chan item)

	go func() { worklist <- item{flag.Args(), 0} }()

	seen := make(map[string]bool)
	for list := range worklist {
		if list.depth > *depth {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				go func(link string, depth int) {
					worklist <- item{crawl(link), depth + 1}
				}(link, list.depth)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Exract(url, false)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
