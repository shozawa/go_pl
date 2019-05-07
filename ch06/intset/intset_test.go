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

func TestHas(t *testing.T) {
	tests := []struct {
		numbers []int
		arg     int
		want    bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1, 2}, 1, true},
		{[]int{65, 66}, 66, true},
		{[]int{1, 2}, 66, false},
	}
	for _, test := range tests {
		set := &IntSet{}
		for _, n := range test.numbers {
			set.Add(n)
		}
		if got := set.Has(test.arg); got != test.want {
			t.Errorf("set.Has(%d) %s is not %v. got=%v", test.arg, set.String(), test.want, got)
		}
	}
}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 2}, []int{3}, "{1 2 3}"},
	}
	for _, test := range tests {
		a := &IntSet{}
		for _, n := range test.a {
			a.Add(n)
		}
		b := &IntSet{}
		for _, n := range test.b {
			b.Add(n)
		}
		a.UnionWith(b)
		if got := a.String(); got != test.want {
			t.Errorf("union is not %s. got=%s", test.want, got)
		}
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{65, 66}, 2},
		{[]int{1, 2, 65, 128}, 4},
	}
	for _, test := range tests {
		set := &IntSet{}
		for _, n := range test.numbers {
			set.Add(n)
		}
		if got := set.Len(); got != test.want {
			t.Errorf("set.Len() %s is not %v. got=%v", set.String(), test.want, got)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		numbers []int
		removed int
		want    string
	}{
		{[]int{1}, 1, "{}"},
		{[]int{1, 2, 3}, 2, "{1 3}"},
		{[]int{1, 2, 3}, 5, "{1 2 3}"},
		{[]int{1, 65, 66}, 65, "{1 66}"},
	}
	for _, test := range tests {
		set := &IntSet{}
		for _, n := range test.numbers {
			set.Add(n)
		}
		set.Remove(test.removed)
		if got := set.String(); got != test.want {
			t.Errorf("set.Remove(%d) %s is not %q. got=%q", test.removed, set.String(), test.want, got)
		}
	}
}
