package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	count := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	for in.Scan() {
		count[in.Text()]++
	}

	fmt.Print("word\tcount\n")

	for _, w := range sortWords(count) {
		fmt.Printf("%s\t%v\n", w, count[w])
	}
}

func sortWords(count map[string]int) []string {
	words := make([]string, 0, len(count))

	for word, _ := range count {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool { return count[words[i]] > count[words[j]] })

	return words
}
