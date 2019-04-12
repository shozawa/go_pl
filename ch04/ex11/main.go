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
		repository := os.Args[2]
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		issue, err := getIssue(repository, id)
		fmt.Print("State\tTitle\tBody\n")
		fmt.Printf("%s\t%s\t%s\n", issue.State, issue.Title, issue.Body)
	case "update":
		repository := os.Args[2]
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		issue, err := getIssue(repository, id)
		if err != nil {
			os.Exit(1)
		}

		serialized, err := json.MarshalIndent(issue, "", "  ")
		if err != nil {
			os.Exit(1)
		}

		tempFile, err := ioutil.TempFile("", "")
		if err != nil {
			os.Exit(1)
		}
		io.Copy(tempFile, strings.NewReader(string(serialized)))
		launchEditor(tempFile.Name())
		defer tempFile.Close()

		editedFile, err := os.Open(tempFile.Name())
		if err != nil {
			os.Exit(1)
		}
		defer editedFile.Close()

		if err := json.NewDecoder(editedFile).Decode(&issue); err != nil {
			os.Exit(1)
		}
		issue.Update()
		fmt.Println(issue)
	case "create":
		repository := os.Args[2]
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
		issue := Issue{
			Repository: repository,
			Title:      title,
			Body:       strings.Join(body, "\n"),
		}
		issue.Create()
	case "close":
		repository := os.Args[2]
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		issue, err := getIssue(repository, id)
		if err != nil {
			os.Exit(1)
		}
		issue.Close()
		issue.Update()
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

type Issue struct {
	Id         int    `json:"-"`
	Repository string `json:"-"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	State      string `json:"state"`
}

func getIssue(repository string, id int) (*Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d", repository, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	issue := Issue{Id: id, Repository: repository}
	if err = json.NewDecoder(res.Body).Decode(&issue); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return &issue, nil
}

func executeCommand(url string, method string, issue *Issue) error {
	token := os.Getenv("GITHUB_PERSONAL_API_TOKEN")
	client := http.DefaultClient

	data, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	return nil
}

func (issue *Issue) URL() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/issues/%d", issue.Repository, issue.Id)
}

func (issue *Issue) Create() {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues", issue.Repository)
	err := executeCommand(url, "POST", issue)
	if err != nil {
		log.Fatal(err)
	}
}

func (issue *Issue) Close() {
	issue.State = "closed"
}

func (issue *Issue) Update() {
	err := executeCommand(issue.URL(), "PATCH", issue)
	if err != nil {
		log.Fatal(err)
	}
}
