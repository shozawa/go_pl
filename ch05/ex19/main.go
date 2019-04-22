package main

import "fmt"

func main() {
	n := panicDouble(10)
	fmt.Println(n) // asert 20
}

func panicDouble(a int) (b int) {
	defer func() {
		b = int(recover().(int))
	}()
	panic(a * 2)
}
