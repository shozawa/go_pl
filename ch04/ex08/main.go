package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

		// TODO: カテゴリ追加
		switch {
		case unicode.IsLetter(r):
			count['L']++
		case unicode.IsNumber(r):
			count['N']++
		case unicode.IsMark(r):
			count['M']++
		}
	}

	fmt.Print("rune\tcount\n")
	// TODO: 出現回数降順で表示する
	for k, v := range count {
		fmt.Printf("%q\t%v\n", k, v)
	}
}
