package strings

import "testing"
import "bufio"

func TestSimpleCase(t *testing.T) {
	reader := NewReader("hello")
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if got := scanner.Text(); got != "hello" {
		t.Errorf("expect hello but got %q", got)
	}
}
