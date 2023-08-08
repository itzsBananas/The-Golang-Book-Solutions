package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Exercise 5.2: Write a function to populate a mapping from element names—p, div, span, and
// so on—to the number of elements with that name in an HTML document tree.

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	index := make(map[string]int, 128)
	for key, val := range visit(index, doc) {
		fmt.Printf("%s: %d\n", key, val)
	}
}

func visit(popl map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return popl
	}
	if n.Type == html.ElementNode {
		popl[n.Data] += 1
	}
	popl = visit(popl, n.FirstChild)
	popl = visit(popl, n.NextSibling)

	return popl
}
