// Exercise 8.5: Take an existing CPU-bound sequential program, such as the Mandelbrot
// program of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its main
// loop in parallel using channels for communication. How much faster does it run on a
// multiprocessor machine? What is the optimal number of goroutines to use?
package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	fmt.Println("NumCPU:", runtime.NumCPU())
	draw()

	for n := 1; n <= runtime.NumCPU(); n++ {
		concurrentDraw(n)
	}
}

func draw() {
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
	// png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func concurrentDraw(num int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(num)

	start := time.Now()
	rows := make(chan int, height)
	go func() {
		for row := 0; row < height; row++ {
			rows <- row
		}
		close(rows)
	}()

	for n := 0; n < num; n++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done. Worker Number:", num, "Used:", time.Since(start))
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

//!-
