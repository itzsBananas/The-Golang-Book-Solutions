// Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those names. Here are two
// example calls:
// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Use only two command line arg's")
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("GET request failed %s", url)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("parsing %s as HTML: %v", url, err)
	}

	elements := ElementsByTagName(doc, os.Args[2:]...)
	for _, element := range elements {
		fmt.Printf("%#v\n", element)
	}
}

func ElementsByTagName(doc *html.Node, tagNames ...string) []*html.Node {
	var elementList []*html.Node

	var forEachNode func(n *html.Node, pre, post func(n *html.Node) bool)
	forEachNode = func(n *html.Node, pre, post func(n *html.Node) bool) {
		if pre != nil && pre(n) {
			elementList = append(elementList, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}
		if post != nil && post(n) {
			elementList = append(elementList, n)
		}
	}

	hasTagName := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, tagName := range tagNames {
			if n.Data == tagName {
				return true
			}
		}
		return false
	}

	forEachNode(doc, hasTagName, nil)
	return elementList
}
