// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"os"
	"strings"
)

// !+
func echo3() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
