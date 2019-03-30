package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}

// [question]
// Slice ではなく配列のポインタを渡す場合は任意の長さの配列を取れない？
func reverse(p *[4]int) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}
