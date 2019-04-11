package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "create":
		tmpfile, _ := ioutil.TempFile("", "")
		defer os.Remove(tmpfile.Name())
		c := exec.Command(getEditor(), tmpfile.Name())
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
		f, _ := os.Open(tmpfile.Name())
		defer f.Close()
		scaner := bufio.NewScanner(f)
		scaner.Scan()
		title := scaner.Text()
		var body []string
		for scaner.Scan() {
			body = append(body, scaner.Text())
		}
		create(title, strings.Join(body, "\n"))
	default:
		fmt.Println("help here.")
	}
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if len(editor) > 1 {
		return editor
	} else {
		return "vi"
	}
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
