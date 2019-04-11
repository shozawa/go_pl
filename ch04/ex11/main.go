package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	create("test", "test body")
}

func create(title string, body string) {
	token := os.Getenv("GITHUB_PERSONAL_API_TOKEN")
	url := "https://api.github.com/repos/shozawa/go_pl/issues"
	client := http.DefaultClient

	json, _ := json.Marshal(map[string]interface{}{
		"title": title,
		"body":  body,
	})

	req, err := http.NewRequest("POST", url, strings.NewReader(string(json)))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+token)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
