package ex16

import "strings"

func Join(sep string, vals ...string) string {
	return strings.Join(vals, sep)
}
