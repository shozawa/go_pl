package tempconv

import "testing"

func TestCToK(t *testing.T) {
	if CToK(Celsius(-273.15)) != Kelvin(0) {
		t.Fatal("dosent't match")
	}
	if KToC(Kelvin(0)) != Celsius(-273.15) {
		t.Fatal("dosent't match")
	}
}
