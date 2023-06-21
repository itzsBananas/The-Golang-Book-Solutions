// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
// bit positions, testing the rightmost bit each time. Compare its performance to the table-
// lookup version.

package main

func shiftCount(x uint64) int {
	var total uint64
	for i := 0; i < 64; i++ {
		total += x & 1
		x >>= 1
	}

	return int(total)
}
