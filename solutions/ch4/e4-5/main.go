package main

import "fmt"

// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
func main() {
	fmt.Println(removeDuplicate([]string{"a", "a", "b", "c", "c", "d", "d"}))
}

func removeDuplicate(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	prevIndex := 0

	for _, chr := range arr[1:] {
		if arr[prevIndex] != chr {
			prevIndex += 1
			arr[prevIndex] = chr
		}
	}
	return arr[:prevIndex+1]
}
