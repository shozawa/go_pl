package intset

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		numbers []int
		want    string
	}{
		{[]int{}, "{}"},
		{[]int{100}, "{100}"},
		{[]int{1, 2, 3}, "{1 2 3}"},
	}
	for _, test := range tests {
		set := &IntSet{}
		for _, n := range test.numbers {
			set.Add(n)
		}
		if got := set.String(); got != test.want {
			t.Errorf("set.String() is not %q. got=%q", test.want, got)
		}
	}
}
