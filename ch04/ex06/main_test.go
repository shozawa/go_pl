package ex06

import "testing"

func TestDeleteSpace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"abcd", "abcd"},
		{"ab cd", "ab cd"},
		{"ab  cd", "ab cd"},
		{"ab   cd", "ab cd"},
		{"abbcd", "abbcd"},
	}
	for _, test := range tests {
		if got := DeleteSpace(test.input); got != test.want {
			t.Errorf("DeleteSpace(%s) = %s, want %s.", test.input, got, test.want)
		}
	}
}
