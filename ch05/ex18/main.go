package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	name, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	// Close以外にエラーが発生している場合はそちらを優先して報告する
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	return
}
