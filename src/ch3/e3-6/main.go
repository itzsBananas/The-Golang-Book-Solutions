// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the
// color value at several points within each pixel and taking the average. The simplest method is
// to divide each pixel into four ‘‘subpixels.’’ Implement it.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	img = supersample(img, width, height)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func supersample(img *image.RGBA, width int, height int) *image.RGBA {
	sampled := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height-1; py += 2 {
		for px := 0; px < width-1; px += 2 {
			top_left, _, _, _ := img.At(px, py).RGBA()
			top_right, _, _, _ := img.At(px+1, py).RGBA()
			bottom_left, _, _, _ := img.At(px, py+1).RGBA()
			bottom_right, _, _, _ := img.At(px+1, py+1).RGBA()
			avg_shade := uint8((top_left + top_right + bottom_left + bottom_right) / 4)
			sampled_color := color.RGBA{avg_shade, avg_shade, avg_shade, 255}

			sampled.Set(px, py, sampled_color)
			sampled.Set(px+1, py, sampled_color)
			sampled.Set(px, py+1, sampled_color)
			sampled.Set(px+1, py+1, sampled_color)
		}
	}
	return sampled
}
