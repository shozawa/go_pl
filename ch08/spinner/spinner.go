package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibn := fib(n)
	fmt.Printf("\rFibonacchi(%d) = %d\n", n, fibn)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
