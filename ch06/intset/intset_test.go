package intset

import "testing"

func TestString(t *testing.T) {
	set := &IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)
	want := "{1 2 3}"
	if got := set.String(); got != want {
		t.Errorf("set.String() is not %q. got=%q", want, got)
	}
}
