package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shozawa/go_pl/ch05/links"
)

func main() {
	worklist := make(chan []string)

	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
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
