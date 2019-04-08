package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"io"
	"os"
	"fmt"
)

type Result struct {
	Poster string
}

// FIXME: ひどい
func main() {
	endpoint := fmt.Sprintf("http://www.omdbapi.com/?apikey=146c9369&t=%s", url.QueryEscape(os.Args[1]))
	res, _ := http.Get(endpoint)
	var result Result
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {

	}
	defer res.Body.Close()

	res, err := http.Get(result.Poster)
	if err != nil {

	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
