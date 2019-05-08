package intset

import "testing"

func TestNew(t *testing.T) {
	set := New(1, 2, 66)
	want := "{1 2 66}"
	if got := set.String(); got != want {
		t.Errorf("New(1, 2, 66) is not %s. got=%s", want, got)
	}
}

func TestAddAll(t *testing.T) {
	set := IntSet{}
	set.AddAll(1, 2, 70, 150)
	if set.String() != "{1 2 70 150}" {
		t.Errorf("set is not {1 2 70 150}. got=%s", set.String())
	}
}

func TestElems(t *testing.T) {
	tests := []struct {
		set  *IntSet
		want []uint64
	}{
		{New(), []uint64{}},
		{New(1, 2, 3), []uint64{1, 2, 3}},
		{New(1, 2, 3, 100), []uint64{1, 2, 3, 100}},
	}
	for _, test := range tests {
		if got := test.set.Elems(); !equalSlice(got, test.want) {
			t.Errorf("set.Elems() is not %v. got=%v", test.want, got)
		}
	}
}

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
		a = a.UnionWith(b)
		if got := a.String(); got != test.want {
			t.Errorf("union is not %s. got=%s", test.want, got)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		a    *IntSet
		b    *IntSet
		want string
	}{
		{New(), New(), "{}"},
		{New(), New(2), "{}"},
		{New(2), New(), "{}"},
		{New(1, 2, 3), New(2), "{2}"},
		{New(1, 2, 3), New(2, 3), "{2 3}"},
		{New(1, 2, 100), New(2, 3), "{2}"},
		{New(2, 3), New(1, 2, 100), "{2}"},
	}
	for _, test := range tests {
		if got := test.a.IntersectWith(test.b); got.String() != test.want {
			t.Errorf("a.IntersectWith(b) is not %s. got=%s", test.want, got)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		a    *IntSet
		b    *IntSet
		want string
	}{
		{New(1), New(), "{1}"},
		{New(1, 2, 3), New(2), "{1 3}"},
		{New(1, 2, 3), New(5), "{1 2 3}"},
		{New(1, 100), New(100), "{1}"},
		{New(1, 2), New(100), "{1 2}"},
		{New(1, 2, 3, 4, 5, 6), New(2, 4, 6, 8), "{1 3 5}"},
		{New(2, 4, 6, 8), New(1, 2, 3, 4, 5, 6), "{8}"},
	}
	for _, test := range tests {
		if got := test.a.DifferenceWith(test.b); got.String() != test.want {
			t.Errorf("a.DifferenceWith(b) is not %s. got=%s", test.want, got)
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

func TestClear(t *testing.T) {
	tests := []struct {
		set  *IntSet
		want *IntSet
	}{
		{New(1, 2, 66), New()},
	}
	for _, test := range tests {
		test.set.Clear()
		testEqual(t, test.want, test.set)
	}
}

func TestCopy(t *testing.T) {
	src := New(1, 2, 3)
	dst := src.Copy()
	if dst.String() != "{1 2 3}" {
		t.Errorf("copy.String() is not {1 2 3}. got=%s", dst.String())
	}

	// mutate
	src.Add(4)

	if dst.String() != "{1 2 3}" {
		t.Errorf("copy.String() is not {1 2 3}. got=%s", dst.String())
	}
}

func testEqual(t *testing.T, want, got *IntSet) bool {
	if want.String() != got.String() {
		t.Errorf("want=%s but got=%s", want.String(), got.String())
		return false
	}
	return true
}

func equalSlice(a, b []uint64) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
