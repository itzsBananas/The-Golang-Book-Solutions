// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
//
// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an
// optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
//
//	func comma(s string) string {
//		n := len(s)
//		if n <= 3 {
//			return s
//		}
//		return comma(s[:n-3]) + "," + s[n-3:]
//	}
func comma(s string) string {
	if len(s) == 0 {
		return ""
	}
	var buf bytes.Buffer

	// strip sign
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	n := strings.Index(s, ".")
	if n <= 3 {
		buf.WriteString(s)
		return buf.String() // change
	}
	base := n % 3

	buf.WriteString(s[0:base])
	if base > 0 {
		buf.WriteByte(',')
	}

	i := base
	for ; i < n-3; i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}
	buf.WriteString(s[i:])

	return buf.String()
}

//!-
