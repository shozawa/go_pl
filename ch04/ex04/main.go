package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := rotate(s1)
	fmt.Println(s1)
	fmt.Println(s2) // expect: [2, 3, 4, 1]
}

func rotate(s []int) []int {
	s = append(s, s[0])
	return s[1:]
}
