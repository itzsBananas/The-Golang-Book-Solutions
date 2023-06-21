package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?
func main() {
	flag.Parse()
	s := flag.Arg(0)
	fmt.Printf("%s -> ", s)
	bytes := []byte(s)
	reverse(bytes)
	fmt.Printf("%s\n", bytes)
}

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func reverse(b []byte) {
	// Reverse each UTF-8 rune
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	// Reverse the whole bytes holding Reversed-UTF-8
	rev(b)
}
