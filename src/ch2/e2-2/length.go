// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no argu-
// ments, and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms, and the like.

package main

import (
	"fmt"
)

type Feet float64
type Meter float64

const (
	FtoMConvert Feet  = 0.3048
	MToFConvert Meter = 3.28084
)

func (f Feet) String() string  { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }

func FtoM(f Feet) Meter { return Meter(f * FtoMConvert) }
func MtoF(m Meter) Feet { return Feet(m * MToFConvert) }
