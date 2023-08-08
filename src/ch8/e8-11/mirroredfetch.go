// Exercise 8.11: Following the approach of mirroredQuery in Section 8.4.4, implement a vari-
// ant of fetch that requests several URLs concurrently. As soon as the first response arrives,
// cancel the other requests.

// func mirroredQuery() string {
// 	responses := make(chan string, 3)
// 	go func() { responses <- request("asia.gopl.io") }()
// 	go func() { responses <- request("europe.gopl.io") }()
// 	go func() { responses <- request("americas.gopl.io") }()
// 	return <-responses // return the quickest response
// 	}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	cancel := make(chan struct{})
	resp := make(chan struct{})

	for _, url := range os.Args[1:] {
		go func(url string) {
			fetch(url, cancel)
			close(cancel)
			fmt.Printf("Got %s\n", url)
			resp <- struct{}{}
		}(url)
	}
	<-resp
	fmt.Println("finished")
}

func fetch(url string, cancel <-chan struct{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		os.Exit(1)
	}

	// Set channel for cancel request.
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// b, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: resp %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("\nfinished %s\n", url)

}
