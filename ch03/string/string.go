package string

import "bytes"

func comma(s string) string {
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
