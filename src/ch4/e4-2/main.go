package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but
// supports a command-line flag to print the SHA384 or SHA512 hash instead.

func main() {
	mode := flag.Int("a", 0, "0: SHA256, 1: SHA384, 2: SHA512")
	flag.Parse()
	str := flag.Arg(0)
	fmt.Printf("Doing hash mode %d on \"%s\"\n", *mode, str)
	switch *mode {
	case 0:
		fmt.Println(sha256.Sum256([]byte(str)))
	case 1:
		fmt.Println(sha512.Sum384([]byte(str)))
	case 2:
		fmt.Println(sha512.Sum512([]byte(str)))
	}
}
