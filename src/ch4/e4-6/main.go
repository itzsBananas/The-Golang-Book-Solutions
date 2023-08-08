package main

import (
	"flag"
	"fmt"
	"unicode"
)

// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
func main() {
	flag.Parse()
	str := flag.Arg(0)
	fmt.Printf("%s\n", squashSpaces([]byte(str)))
}

func squashSpaces(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	nextPlacement, prevChar := 0, 0

	for i := 0; i < len(s); {
		if unicode.IsSpace(rune(s[i])) {
			addedLength := copy(s[nextPlacement:], s[prevChar:i])
			nextPlacement += addedLength
			for j := i + 1; j < len(s); j++ {
				if !unicode.IsSpace(rune(s[j])) {
					prevChar, i = j-1, j
					break
				}
			}
		}
		i++

	}
	addedLength := copy(s[nextPlacement:], s[prevChar:])
	return s[:nextPlacement+addedLength+1]
}
