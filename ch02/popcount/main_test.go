package popcount

import "testing"

func TestPopcount(t *testing.T) {
	counts := [...]int{
		Count1(18446744073709551615),
		Count2(18446744073709551615),
		Count3(18446744073709551615),
		Count4(18446744073709551615),
		Count5(18446744073709551615),
	}

	for i, count := range counts {
		if count != 64 {
			t.Fatalf("expect 64 but got %d\t i:%d\n", count, i)
		}
	}
}

func BenchmarkPopcount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count1(18446744073709551615)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count2(18446744073709551615)
	}
}

func BenchmarkPopcount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count3(18446744073709551615)
	}
}

func BenchmarkPopcount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count4(18446744073709551615)
	}
}

func BenchmarkPopcount5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count5(18446744073709551615)
	}
}
