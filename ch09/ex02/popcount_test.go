package popcount

import (
	"testing"
)

func TestCount(t *testing.T) {
	if got := Count(100); got != 3 {
		t.Errorf("Count(100) is not 3. got=%d", got)
	}
}
