package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
)

// Rune の数を数える
func main() {
	input := bufio.NewReader(os.Stdin)
	count := make(map[rune]int)

	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}

		switch {
		case unicode.IsLetter(r):
			count['L']++
		case unicode.IsNumber(r):
			count['N']++
		case unicode.IsMark(r):
			count['M']++
		}
	}

	runes := []rune{}

	for r, _ := range count {
		runes = append(runes, r)
	}

	sort.Slice(runes, func(i, j int) bool { return count[runes[i]] > count[runes[j]] })

	fmt.Print("rune\tcount\n")
	for _, r := range runes {
		fmt.Printf("%q\t%v\n", r, count[r])
	}
}
