package main

import (
	"fmt"
	"testing"
)

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(12, 24, 36)
	fmt.Println(x.String()) // {12, 24, 36}

	x.AddAll(6, 13, 25)
	fmt.Println(x.String()) // {6, 12, 13, 24, 25, 36}
}
