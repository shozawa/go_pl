package main

import (
	"fmt"
)

const (
	UNCHECKED = iota
	TEMPORARY
	PERMANENT
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func toposort(m map[string][]string) []string {
	var orderd []string
	var visit func(string)
	seen := make(map[string]int)

	visit = func(item string) {
		if seen[item] == TEMPORARY {
			panic("not DAG")
		} else if seen[item] == UNCHECKED {
			seen[item] = TEMPORARY
			for _, dependency := range m[item] {
				visit(dependency)
			}
			seen[item] = PERMANENT
			orderd = append(orderd, item)
		}
	}
	for item, _ := range m {
		visit(item)
	}
	return orderd
}

func isSorted(m map[string][]string, sorted []string) bool {
	for i, item := range sorted {
		for _, dependency := range m[item] {
			order := findIndex(sorted, dependency)
			if order == -1 {
				// リストの中に依存関係がある項目が存在しない
				return false
			}
			if order > i {
				return false
			} 
		}
	}
	return true 
}

func findIndex(list []string, s string) int {
	for i, _ := range list {
		if list[i] == s {
			return i 
		}
	}
	return -1
}