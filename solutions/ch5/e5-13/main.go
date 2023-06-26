// Exercise 5.13: Modify crawl to make local copies of the pages it finds, creating directories as
// necessary. Don’t make copies of pages that come from a different domain. For example, if the
// original page comes from golang.org, save all files from there, but exclude ones from
// vimeo.com.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

func savePage(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("saved/%s.txt", strings.ReplaceAll(url, "/", ".")))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	resp.Write(f)

	return nil
}

func findHostName(input string) string {
	url, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	trimmed := strings.TrimPrefix(url.Hostname(), "www.")
	trimmed = strings.TrimSuffix(trimmed, "/")
	return trimmed
}

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	allowedSites := make(map[string]bool)
	for _, input := range worklist {
		allowedSites[findHostName(input)] = true
	}
	fmt.Println(allowedSites)

	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			hostName := findHostName(item)
			if !seen[item] && allowedSites[hostName] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

// !+crawl
func crawl(url string) []string {
	// fmt.Println(url)
	savePage(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
