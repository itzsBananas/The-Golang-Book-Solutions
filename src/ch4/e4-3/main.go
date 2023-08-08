package main

import "fmt"

// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// func reverse(s []int) {
// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 	s[i], s[j] = s[j], s[i]
// }
// }
