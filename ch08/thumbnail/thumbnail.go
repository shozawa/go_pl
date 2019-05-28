package main

import (
	"log"
	"path/filepath"

	"github.com/adonovan/gopl.io/ch8/thumbnail"
)

func main() {
	ch := make(chan struct{})
	files, err := filepath.Glob("./ch08/images/*.jpg")
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(file)
	}
	for range files {
		<-ch
	}
}
