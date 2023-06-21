// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
// of PopCount that counts bits by using this fact, and assess its performance.

package main

func shiftCount(x uint64) int {
	i := 0
	for x != 0 {
		x = x & (x - 1)
		i += 1
	}

	return i
}
