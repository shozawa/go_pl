package counter

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"one", 1},
		{"one two", 2},
		{"one two three", 3},
		{"", 0},
		{"one\ntwo\n", 2},
	}
	for _, test := range tests {
		var counter WordCount
		fmt.Fprint(&counter, test.input)
		if got := int(counter); got != test.want {
			t.Errorf("count of %q is not %d. got=%d", test.input, test.want, got)
		}
	}
}

func TestLineCount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"one", 1},
		{"one two", 1},
		{"", 0},
		{"one\ntwo\n", 2},
	}
	for _, test := range tests {
		var counter LineCount
		fmt.Fprint(&counter, test.input)
		if got := int(counter); got != test.want {
			t.Errorf("count of %q is not %d. got=%d", test.input, test.want, got)
		}
	}
}
