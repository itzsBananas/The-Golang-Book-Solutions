package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(variadicStringJoin(", ", "Hello", "World"))
	fmt.Println(variadicStringJoin(", "))
	fmt.Println(variadicStringJoin(", ", "Dog"))
}

func variadicStringJoin(sep string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(strings[0])
	for _, str := range strings[1:] {
		if len(str) != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(str)
	}
	return buf.String()
}
