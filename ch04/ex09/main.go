package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	for in.Scan() {
		count[in.Text()]++
	}

	fmt.Print("word\tcount\n")

	// TODO: 出現回数降順で並び替え
	for w, c := range count {
		fmt.Printf("%s\t%v\n", w, c)
	}
}
