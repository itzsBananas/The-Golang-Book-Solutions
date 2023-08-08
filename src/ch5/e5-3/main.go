// Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into <script> or <style> elements, since their contents are not visible
// in a web browser.

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, val := range visit(nil, doc) {
		fmt.Printf("%s\n", val)
	}

}

func visit(str []string, n *html.Node) []string {
	if n == nil {
		return str
	}
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		for _, line := range strings.Split(n.Data, "\n") {
			if len(line) != 0 {
				str = append(str, line)
			}
		}
	}
	str = visit(str, n.FirstChild)

	return visit(str, n.NextSibling)
}
