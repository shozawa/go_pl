package main

import (
	"fmt"
	"log"
	"os"
	"github.com/shozawa/go_pl/ch05/graph"
	"github.com/shozawa/go_pl/ch05/links"
)

func main() {
	graph.BFS(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Exract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}