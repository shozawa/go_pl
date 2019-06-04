package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizesMap := make(map[string]chan int64)
	for _, root := range roots {
		fileSizesMap[root] = make(chan int64)
	}

	// Question: 単一の goroutine の中でループを回したら deadlock になった。なぜ？
	/*
		go func() {
			for _, root := range roots {
				walkDir(root, fileSizesMap[root])
				close(fileSizesMap[root])
			}
		}()
	*/

	for _, root := range roots {
		go func(root string) {
			walkDir(root, fileSizesMap[root])
			close(fileSizesMap[root])
		}(root)
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	for root, fileSizes := range fileSizesMap {
		var nfiles, nbytes int64
	loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(root, nfiles, nbytes)
			}
		}
		printDiskUsage(root, nfiles, nbytes)
	}
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s\t%d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v'n", err)
		return nil
	}
	return entries
}
