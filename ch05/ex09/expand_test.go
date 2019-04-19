package expand

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"foo $bar", "foo BAR"},
		{"foo $bar buzz", "foo BAR buzz"},
		{"foo $bar $buzz", "foo BAR BUZZ"},
		{"foo", "foo"},
		{"", ""},
	}
	for _, test := range tests {
		if got := Expand(test.input, strings.ToUpper); got != test.want {
			t.Errorf("want %v but got %v.", test.want, got)
		}
	}
}
