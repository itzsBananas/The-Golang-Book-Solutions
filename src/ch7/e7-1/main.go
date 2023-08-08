// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var c WordCounter
	c.Write([]byte("hello world"))
	fmt.Println(c)

	var name = "Dolly the sheep"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var d LineCounter
	d.Write([]byte("hello, it's \n charles xavier"))
	fmt.Println(d)

	fmt.Fprintf(&d, "hello, \n asdf \n bob %s", name)
	fmt.Println(d)
}

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	length := len(strings.Split(string(p), " "))
	*c += WordCounter(length) // convert int to ByteCounter
	return length, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	length := len(strings.Split(string(p), "\n"))
	*c += LineCounter(length) // convert int to ByteCounter
	return length, nil
}
