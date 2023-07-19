// Exercise 7.2: Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cw, value := CountingWriter(os.Stdout)
	fmt.Println(*value)
	cw.Write([]byte("chicken "))
	fmt.Println()
	fmt.Println(*value)
	cw.Write([]byte("dinner"))
	fmt.Println()
	fmt.Println(*value)
}

type countingWriterStruct struct {
	writer  io.Writer
	counter int64
}

func (i *countingWriterStruct) Write(p []byte) (n int, err error) {
	n, err = i.writer.Write(p)
	i.counter += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriterStruct{w, 0}
	return cw, &cw.counter
}
