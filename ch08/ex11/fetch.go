package main

import (
	"context"
	"fmt"
	"net/http"
)

var status = make(chan string)

// FXIME: 間違ってそう
func main() {
	urls := []string{
		"https://duckduckgo.com",
		"https://www.bing.com",
		"https://www.google.com",
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	fetch(ctx, urls)
	fmt.Println(<-status)
	cancelFunc()
}

func fetch(ctx context.Context, urls []string) {
	for _, url := range urls {
		go func(url string) {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Print(err)
			}
			req = req.WithContext(ctx)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Print(err)
			}
			defer res.Body.Close()
			fmt.Println(req)
			status <- res.Status
		}(url)
	}
}
