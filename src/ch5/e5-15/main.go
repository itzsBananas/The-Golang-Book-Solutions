// Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these
// functions do when called with no arguments? Write variants that require at least one argu-
// ment.

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max())
}

func max(nums ...int) int {
	if len(nums) == 0 {
		log.Fatal("There must be at least one function argument")
	}
	maxNum := nums[0]
	for _, num := range nums[1:] {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func min(nums ...int) int {
	if len(nums) == 0 {
		log.Fatal("There must be at least one function argument")
	}
	minNum := nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
