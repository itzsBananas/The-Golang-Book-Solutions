// Exercise 5.8: Modify forEachNode so that the pre and post functions return a boolean result
// indicating whether to continue the traversal. Use it to write a function ElementByID with the
// following signature that finds the first HTML element with the specified id attribute. The
// function should stop the traversal as soon as a match is found.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	const (
		url = "https://golang.org"
		id  = "quote_slide1"
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("getting %s: %s", url, resp.Status))
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("parsing %s as HTML: %v", url, err))
	}

	found := ElementByID(doc, id)
	fmt.Printf("%#v\n", found)
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (dest *html.Node) {
	if pre != nil && pre(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dest = forEachNode(c, pre, post)
	}
	if post != nil && post(n) {
		return n
	}
	return dest
}

func ElementByID(doc *html.Node, id string) *html.Node {
	continueTraversal := func(doc *html.Node) bool {
		if doc.Type == html.ElementNode {
			for _, a := range doc.Attr {
				if a.Key == "id" && a.Val == id {
					return false
				}
			}
		}
		return true
	}
	return forEachNode(doc, continueTraversal, nil)
}
