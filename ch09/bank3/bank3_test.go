package bank2

import "testing"

func TestSimpleCase(t *testing.T) {
	if got := Balance(); got != 0 {
		t.Errorf("Balance() is not 0. got=%d.", got)
	}
	Deposit(100)
	if got := Balance(); got != 100 {
		t.Errorf("Balance() is not 100. got=%d.", got)
	}
}
