package string

import (
	"bytes"
	"sort"
	"strings"
)

func Comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i, r := range s {
		digit := n - i
		buf.WriteRune(r)
		if digit != 1 && digit%3 == 1 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

func Sort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func IsAnagram(a, b string) bool {
	a = Sort(strings.ToLower(strings.Replace(a, " ", "", -1)))
	b = Sort(strings.ToLower(strings.Replace(b, " ", "", -1)))
	return a == b
}
