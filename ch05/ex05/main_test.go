package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	input := "this is a pen"
	want := 4
	if got := countWords(input); got != want {
		t.Errorf("input: %v\texpect %v but got %v", input, want, got)
	}
}
