package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func execDB(db *sql.DB, q string) {
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./xkcd.db")
	if err != nil {
		log.Fatal(err)
	}

	q := `
        CREATE TABLE comics (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
		  title VARCHAR(255) NOT NULL,
		  url VARCHAR(255) NOT NULL,
		  body TEXT
          created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
        )
    `
	execDB(db, q)
	defer db.Close()
}
