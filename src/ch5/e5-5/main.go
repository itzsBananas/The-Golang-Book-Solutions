// Exercise 5.5: Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)

package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImages("https://golang.org"))
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	switch n.Type {
	case html.TextNode:
		list := strings.Split(strings.ReplaceAll(n.Data, "\n", " "), " ")
		words += len(list)
	case html.ElementNode:
		if n.Data == "img" {
			images += 1
		}
	}

	w1, i1 := countWordsAndImages(n.FirstChild)
	w2, i2 := countWordsAndImages(n.NextSibling)
	words += w1 + w2
	images += i1 + i2

	return
}
