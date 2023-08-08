package main

import "fmt"

// Exercise 4.4: Write a version of rotate that operates in a single pass.
func main() {
	fmt.Println(rotateLeft([]int{1, 2, 3, 4}, 2))
}

// obvious way
func rotateLeft(s []int, place int) []int {
	n := len(s)
	rotated := make([]int, n, cap(s))
	for i := 0; i < n; i++ {
		rotated[i] = s[(i+place)%n]
	}

	return rotated
}
