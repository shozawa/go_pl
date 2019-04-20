package main

import (
	"fmt"
	"sort"
)

const (
	UNCHECKED = iota
	TEMPORARY
	PERMANENT
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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

	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, item := range keys {
		visit(item)
	}

	return orderd
}
