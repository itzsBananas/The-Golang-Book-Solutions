// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error occurred when read file")
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	freq := make(map[string]int, 64)
	for scanner.Scan() {
		freq[scanner.Text()]++
	}

	fmt.Println(freq)
}
