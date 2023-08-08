package main

import "fmt"

// Exercise 5.19: Use panic and recover to write a function that contains no return statement
// yet returns a non-zero value.

func main() {
	fmt.Println(something())
}

func something() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = 5
		}
	}()
	panic("lol")
}
