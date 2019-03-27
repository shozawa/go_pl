package string

import (
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"100", "100"},
		{"1000", "1,000"},
		{"100000", "100,000"},
		{"1000000", "1,000,000"},
	}
	for _, test := range tests {
		if got := Comma(test.input); got != test.want {
			t.Errorf("Comma(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestSort(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"abcde", "abcde"},
		{"edcba", "abcde"},
	}
	for _, test := range tests {
		if got := Sort(test.input); got != test.want {
			t.Errorf("Sort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		input [2]string
		want  bool
	}{
		{[2]string{"not", "anagram"}, false},
		{[2]string{"astronomers", "nomorestars"}, true},
		{[2]string{"Statue of Liberty", "built to stay free"}, true},
	}
	for _, test := range tests {
		if got := IsAnagram(test.input[0], test.input[1]); got != test.want {
			t.Errorf("IsAnagram(%s, %s) = %t, want %t", test.input[0], test.input[1], got, test.want)
		}
	}
}
