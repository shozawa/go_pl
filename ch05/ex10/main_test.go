package main

import "testing"

func TestIsSorted(t *testing.T) {
	courses := map[string][]string{
		"algorithms":      {"data structures"},
		"data structures": {"discrete math"},
	}
	sorted := []string{
		"discrete math",
		"data structures",
		"algorithms",
	}
	unsorted := []string{
		"discrete math",
		"algorithms",
		"data structures",
	}
	tests := []struct {
		graph map[string][]string
		list  []string
		want  bool
	}{
		{courses, sorted, true},
		{courses, unsorted, false},
	}
	for _, test := range tests {
		if got := isSorted(test.graph, test.list); got != test.want {
			t.Errorf("graph\t%v\tsorted\t%vwant%v, but %v", test.graph, test.list, test.want, got)
		}

	}
}
