package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "get":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		issue, err := getIssue(id)
		fmt.Print("State\tTitle\tBody\n")
		fmt.Printf("%s\t%s\t%s\n", issue.State, issue.Title, issue.Body)
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		issue, err := getIssue(id)
		j, _ := json.MarshalIndent(issue, "", "  ")
		tempFile, _ := ioutil.TempFile("", "")
		io.Copy(tempFile, strings.NewReader(string(j)))
		launchEditor(tempFile.Name())
		defer tempFile.Close()

		f, err := os.Open(tempFile.Name())
		if err != nil {
			os.Exit(1)
		}
		defer f.Close()

		if err := json.NewDecoder(f).Decode(&issue); err != nil {
			os.Exit(1)
		}

		update(id, issue)
		fmt.Println(issue)
	case "create":
		tmpfile, _ := ioutil.TempFile("", "")
		defer os.Remove(tmpfile.Name())

		launchEditor(tmpfile.Name())

		f, err := os.Open(tmpfile.Name())
		if err != nil {
			os.Exit(1)
		}
		defer f.Close()

		scaner := bufio.NewScanner(f)
		scaner.Scan()
		title := scaner.Text()
		var body []string
		for scaner.Scan() {
			body = append(body, scaner.Text())
		}
		create(title, strings.Join(body, "\n"))
	case "close":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		closeIssue(id)
	default:
		fmt.Println("help here.")
	}
}

func launchEditor(filename string) {
	c := exec.Command(getEditor(), filename)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if len(editor) > 1 {
		return editor
	} else {
		return "vi"
	}
}

func closeIssue(id int) {
	token := os.Getenv("GITHUB_PERSONAL_API_TOKEN")
	url := fmt.Sprintf("https://api.github.com/repos/shozawa/go_pl/issues/%d", id)

	client := http.DefaultClient

	json, _ := json.Marshal(map[string]interface{}{
		"state": "closed",
	})

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(json)))

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

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

func getIssue(id int) (*Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/shozawa/go_pl/issues/%d", id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err = json.NewDecoder(res.Body).Decode(&issue); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return &issue, nil
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

func update(id int, issue *Issue) {
	token := os.Getenv("GITHUB_PERSONAL_API_TOKEN")
	url := fmt.Sprintf("https://api.github.com/repos/shozawa/go_pl/issues/%d", id)
	client := http.DefaultClient

	json, _ := json.Marshal(issue)

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(json)))

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
