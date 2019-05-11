package treesort

import "testing"

func TestTreeSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		Sort(test.input)
		assertEqual(t, test.want, test.input)
	}
}

func assertEqual(t *testing.T, want, got []int) bool {
	if want == nil && got == nil {
		return true
	}
	if len(want) != len(got) {
		panic("len(want) is not len(got)")
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want=%v. got=%v", want, got)
			return false
		}
	}
	return true
}
