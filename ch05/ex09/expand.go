package expand

import "strings"

func Expand(s string, f func(string) string) string {
	split := strings.SplitN(s, "$", 2)
	if len(split) != 2 {
		return s
	}
	return split[0] + f(split[1])
}
