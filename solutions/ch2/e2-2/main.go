// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no argu-
// ments, and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms, and the like.

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			readFloat(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			readFloat(t)
		}
	}
}

func readFloat(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))

	feet := Feet(t)
	meter := Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		feet, FtoM(feet), meter, MtoF(meter))
}

//!-
