package comic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Result struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Img        string
	Transcript string
	Title      string
	Alt        string
	Day        string
}

func (result *Result) Save(db *sql.DB) {
	_, err := db.Exec(`
	INSERT INTO comics(id, title, url, body) VALUES (?, ?, ?, ?)
	`, result.Num, result.Title, result.Img, result.Transcript)
	if err != nil {
		log.Fatal(err)
	}
}

func Fetch(id int) (*Result, error) {
	res, _ := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", id))
	var result Result
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return &result, nil
}

func SearchByKeyword(word string, db *sql.DB) (*sql.Rows, error) {
	// TODO: エスケープ処理
	word = "%" + word + "%"
	rows, err := db.Query("SELECT title, url, body FROM comics WHERE title LIKE ? OR body LIKE ?", word, word)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
