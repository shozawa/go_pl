package ex15

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{3, 2}, 3},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for _, test := range tests {
		if got := Max(test.input...); got != test.want {
			t.Errorf("input %v want %d but got %d\n", test.input, test.want, got)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{3, 2}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for _, test := range tests {
		if got := Min(test.input...); got != test.want {
			t.Errorf("input %v want %d but got %d\n", test.input, test.want, got)
		}
	}
}
