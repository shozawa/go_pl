package main

import (
	"fmt"
	"net/http"
	"sync"
)

var status = make(chan string)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://duckduckgo.com",
		"https://www.bing.com",
	}
	fetch(urls)
	for s := range status {
		fmt.Println(s)
	}
}

func fetch(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				fmt.Print(err)
			}
			status <- res.Status
		}(url)
	}
	// Question: クローザーは goroutine じゃなきゃいけない？
	// goroutine の外に出すとハングする
	go func() {
		wg.Wait()
		close(status)
	}()
}
