package expand

import (
	"strings"
)

func Expand(s string, mapping func(string) string) string {
	var expanded []string
	// 余分な空白は削除されるけど...いいか
	for _, word := range strings.Fields(s) {
		if strings.HasPrefix(word, "$") {
			word = mapping(word[1:])
		}
		expanded = append(expanded, word)
	}
	return strings.Join(expanded, " ")
}
