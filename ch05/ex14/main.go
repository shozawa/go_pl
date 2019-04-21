package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shozawa/go_pl/ch05/graph"
)

// コマンドライン引数で与えられたディレクトリをパースして出力する
func main() {
	graph.BFS(crawl, os.Args[1:])
}

func crawl(path string) (list []string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	// [要確認] path がディレクトリではなかった場合エラーになってくれるはず？
	if err != nil {
		return
	}
	for _, file := range files {
		list = append(list, filepath.Join(path, file.Name()))
	}
	return
}
