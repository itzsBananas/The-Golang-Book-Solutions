// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
// Exercise 3.4: Following the approach of the Lissajous example in Section 1.7, construct a web
// server that computes surfaces and writes SVG data to the client. The server must set the Con-
// tent-Type header like this:
// Allow the client to specify values like height, width, and color as
// HTTP request parameters.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	color         = "white"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		heightParam := r.URL.Query().Get("height")
		if len(heightParam) == 0 {
			heightParam = strconv.FormatInt(height, 10)
		}

		widthParam := r.URL.Query().Get("width")
		if len(widthParam) == 0 {
			widthParam = strconv.FormatInt(width, 10)
		}

		colorParam := r.URL.Query().Get("color")
		if len(colorParam) == 0 {
			colorParam = color
		}
		fmt.Printf("height: %s\t width: %s\t color: %s\n", heightParam, widthParam, colorParam)

		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
			"width='%s' height='%s'>", colorParam, widthParam, heightParam)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
