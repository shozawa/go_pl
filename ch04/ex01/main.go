package main 

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	fmt.Println(countDiff(a, b))
}

func countDiff(a, b [32]byte) int {
	count := 0

	for i, _ := range a {
		count += popcount(a[i] ^ b[i])
	}

	return count
}

func popcount(x byte) int {
	count := 0
	for ; x != 0; x &= x - 1 {
		count++
	}
	return count
}
