// Exercise 5.9: Write a function expand(s string, f func(string) string) string that
// replaces each substring ‘‘$foo’’ within s by the text returned by f("foo").

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	substring       = "$foo"
	substringLength = len(substring)
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(expand(input.Text(), funny))
	}
}

func expand(s string, f func(string) string) string {
	var buf bytes.Buffer

	left := strings.Index(s, substring)
	for left != -1 {
		right := left + substringLength

		buf.WriteString(s[:left])
		buf.WriteString(f(s[left:right]))

		s = s[right:]
		left = strings.Index(s, substring)
	}
	buf.WriteString(s)

	return buf.String()
}

func funny(s string) string {
	if len(s) < 2 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteRune(rune(s[0]))
	for _, char := range s[1:] {
		buf.WriteString(fmt.Sprintf("-%s", string(char)))
	}

	return buf.String()
}
