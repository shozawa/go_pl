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

	// FIXME: 全部取る。あとRetry入れる。404を返すまで取得とか？
	for i := 1; i < 20; i++ {
		result, err := comic.Fetch(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result.Save(db)
	}
}
