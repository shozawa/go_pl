package ex05

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"abcd", "abcd"},
		{"abbcd", "abcd"},
		{"abbbcd", "abcd"},
		{"ababcd", "ababcd"},
	}
	for _, test := range tests {
		if got := RemoveDuplicate(test.input); got != test.want {
			t.Errorf("RemoveDuplicate(%s) = %s, want %s.", test.input, got, test.want)
		}
	}
}
