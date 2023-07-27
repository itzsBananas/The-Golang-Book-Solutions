// Exercise 8.10: HTTP requests may be cancelled by closing the optional Cancel channel in the
// http.Request struct. Modify the web crawler of Section 8.6 to support cancellation.

// Hint: the http.Get convenience function does not give you an opportunity to customize a
// Request. Instead, create the request using http.NewRequest, set its Cancel field, then per-
// form the request by calling http.DefaultClient.Do(req).
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	list, err := Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string, cancel <-chan struct{}) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}

	// Set channel for cancel request.
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// !+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	cancel := make(chan struct{})

	var wg sync.WaitGroup

	// cancelled := func() bool {
	// 	select {
	// 	case <-cancel:
	// 		return true
	// 	default:
	// 		return false
	// 	}
	// }

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Println("closed :O")
		close(cancel)
	}()

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(num int) {
			func() {
				for {
					select {
					case link := <-unseenLinks:
						foundLinks := crawl(link, cancel)
						go func() {
							select {
							case <-cancel: // Cancel signal.
							default:
								worklist <- foundLinks
							}
						}()
					case <-cancel:
						return
					}
				}
			}()
			fmt.Printf("Canceled in crawler %d\n", num)
			wg.Done()
		}(i)
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for {
		select {
		case <-cancel:
			panic("smile")
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
}

//!-
