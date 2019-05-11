package counter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{"", 0},
		{"hello, ", 7},
		{"world", 12},
		{".", 13},
	}
	var b bytes.Buffer
	w, count := CountingWriter(&b)
	// want のカウントは累積していく
	for _, test := range tests {
		fmt.Fprint(w, test.input)
		if *count != test.want {
			t.Errorf("total count of %q is not %d. got=%d", b.String(), test.want, *count)
		}
	}
}
