// Exercise 10.1: Extend the jpeg program so that it converts any supported input format to any
// output format, using image.Decode to detect the input format and a flag to select the output
// format.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	imgFormat := flag.String("f", "png", "Select Image Format...")

	if err := toOutput(os.Stdin, os.Stdout, *imgFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toOutput(in io.Reader, out io.Writer, imgFormat string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format := strings.ToLower(imgFormat); format {
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	case "png":
		return png.Encode(out, img)
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("wrong Format Specified (%s)", format)
	}
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
