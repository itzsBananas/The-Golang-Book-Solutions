package main

import (
	"fmt"
	"time"
)

func main() {
	benchmark(echo1)
	benchmark(echo2)
	benchmark(echo3)
}

func benchmark(a func()) {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		a()
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
