package main

import (
	"fmt"
	"runtime"
	"time"
)
/* 32GiB Memory
 * -------------------------
 * num goroutine: 14000002
 * time: 35.384445465s
 * -------------------------
 * num goroutine: 17000002
 * time 42.568388708s
 */

func main() {
	last := make(chan struct{})
	in := pipe(last)
	for i := 0; i < 14000000; i++ {
		in = pipe(in)
	}

	start := time.Now()

	fmt.Println(runtime.NumGoroutine())

	in <- struct{}{}
	<-last

	fmt.Println(time.Since(start))
}

func pipe(out chan<- struct{}) chan struct{} {
	in := make(chan struct{})
	go func() {
		<-in
		out <- struct{}{}
	}()

	return in
}
