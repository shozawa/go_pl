package ex16

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"A", "B"}, "A B"},
		{[]string{"hello", "go", "lang"}, "hello go lang"},
		{[]string{"A"}, "A"},
		{[]string{""}, ""},
	}
	for _, test := range tests {
		if got := Join(" ", test.input...); got != test.want {
			t.Errorf("input %v want %s but got %s.\n", test.input, test.want, got)
		}
	}
}
