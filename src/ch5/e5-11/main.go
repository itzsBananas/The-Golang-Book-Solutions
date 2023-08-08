// Exercise 5.11: The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
	"log"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
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

//!-table

// !+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	// var keys []string
	for key := range m {
		visitAll([]string{key})
	}

	index := make(map[string]int, len(order))
	for i, name := range order {
		index[name] = i + 1
	}

	for class, class_prereqs := range prereqs {
		for _, preq := range class_prereqs {
			if index[class] < index[preq] {
				log.Fatalf("The class (%s) cannot be taken without taking the class (%s)\n", class, preq)
			}
		}
	}

	// sort.Strings(keys)
	// visitAll(keys)
	return order
}

// !+ Checks if c1 is a prequisite of c2

//!-main
