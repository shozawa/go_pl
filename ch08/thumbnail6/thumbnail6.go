package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/adonovan/gopl.io/ch8/thumbnail"
)

func main() {
	images := make(chan string)

	go func() {
		files, err := filepath.Glob("./ch08/images/*.jpg")
		if err != nil {
			log.Println(err)
		}
		for _, file := range files {
			images <- file
		}
		close(images)
	}()

	size := makeThumbnail(images)
	fmt.Println(size)
}

func makeThumbnail(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
