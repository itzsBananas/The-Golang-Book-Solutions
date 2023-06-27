// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"testing"
)

func Test_Example_one(t *testing.T) {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // 3

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // 2

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())    // 4

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	x.Remove(42)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // 3

	x.Remove(7)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // 3

	x.Remove(9)
	fmt.Println(x.String()) // "{1 144}"
	fmt.Println(x.Len())    // 2

	x.Clear()
	fmt.Println(x.String()) // "{}"
	fmt.Println(x.Len())    // 0
	//!-main

	// Output:
	// {1 9 144}
	// 3
	// {9 42}
	// 2
	// {1 9 42 144}
	// 4
	// true false
}

func Test_Example_two(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
