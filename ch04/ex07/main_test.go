package ex07

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("abcd"), []byte("dcba")},
	}
	for _, test := range tests {
		if Reverse(test.input); string(test.input) != string(test.want) {
			t.Errorf("%v != %v", test.input, test.want)
		}
	}
}
