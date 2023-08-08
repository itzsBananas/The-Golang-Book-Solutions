package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid # of args")
		return
	}
	a, b := os.Args[1], os.Args[2]

	fmt.Printf("%s is an anagram of %s\t %t\n", a, b, isAnagram(a, b))
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}

	const (
		bias = 'a'
	)
	var count [26]int
	for _, ch1 := range s1 {
		count[ch1-bias]++
	}

	for _, ch2 := range s2 {
		index := ch2 - bias
		count[index]--
		if count[index] < 0 {
			return false
		}
	}

	return true
}
