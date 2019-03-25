package string

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"100", "100"},
		{"1000", "1,000"},
		{"1000000", "1,000,000"},
	}
	for _, test := range tests {
		if got := comma(test.input); got != test.want {
			t.Errorf("comma(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
