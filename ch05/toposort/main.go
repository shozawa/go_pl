package main

import (
	"fmt"
	"sort"
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
	seen := make(map[string]bool)

	visit = func(item string) {
		if seen[item] {
			return
		}
		seen[item] = true
		for _, dependency := range m[item] {
			visit(dependency)
		}
		orderd = append(orderd, item)
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
