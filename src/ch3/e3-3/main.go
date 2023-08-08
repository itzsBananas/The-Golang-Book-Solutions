// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.

// Exercise 3.3: Color each polygon based on its height, so that the peaks are colored red
// (#ff0000) and the valleys blue (#0000ff).
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z0, err0 := corner(i+1, j)
			bx, by, z1, err1 := corner(i, j)
			cx, cy, z2, err2 := corner(i, j+1)
			dx, dy, z3, err3 := corner(i+1, j+1)
			if err0 || err1 || err2 || err3 {
				continue
			}
			max_z := math.Max(z0, math.Max(z1, math.Max(z2, z3)))
			fill := 0xffffff
			if max_z > 0 {
				fill = 0xff0000
			} else {
				fill = 0x0000ff
			}

			fmt.Printf("<polygon fill=\"#%06x\" points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				fill, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)
	if err {
		return 0, 0, 0, true
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, false
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r
	if math.IsNaN(r) {
		return 0, true
	}
	return r, false
}

//!-
