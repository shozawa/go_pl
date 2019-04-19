package expand

import (
	"bufio"
	"strings"
)

func Expand(s string, f func(string) string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	var expanded []string

	for scanner.Scan() {
		word := scanner.Text()
		if strings.HasPrefix(word, "$") {
			word = strings.ReplaceAll(word, "$", "")
			word = f(word)
		}
		expanded = append(expanded, word)
	}
	return strings.Join(expanded, " ")
}
