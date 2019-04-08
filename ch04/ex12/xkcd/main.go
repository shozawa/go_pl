package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shozawa/go_pl/ch04/ex12/comic"
)

func main() {
	db, err := sql.Open("sqlite3", "./xkcd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	keyword := os.Args[1]
	rows, _ := comic.SearchByKeyword(keyword, db)

	for rows.Next() {
		var title string
		var url string
		var body string
		rows.Scan(&title, &url, &body)
		fmt.Printf("%s\t%s\t%s\n", title, url, body)
	}
}
