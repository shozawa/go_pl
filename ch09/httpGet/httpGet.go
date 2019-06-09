package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/shozawa/go_pl/ch09/memo5"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
		"https://golang.org",
	}

	m := memo5.New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
