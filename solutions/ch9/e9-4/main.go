// Exercise 9.4: Construct a pipeline that connects an arbitrary number of goroutines with chan-
// nels. What is the maximum number of pipeline stages you can create without running out of
// memory? How long does a value take to transit the entire pipeline?

package main

import (
	"fmt"
	"time"
)

// 9999999 too many
// 1000000 goroutines
// 8.67s elapsed
func main() {
	start := time.Now()
	defer func() { fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds()) }()

	n := 9999999
	last := make(chan int)
	prev := last
	var next chan int
	for i := 0; i < n; i++ {
		next = make(chan int)
		go func(p, n chan int) {
			p <- (<-n + 1)
		}(prev, next)
		prev = next
	}
	next <- 0
	fmt.Println(<-last)
}
