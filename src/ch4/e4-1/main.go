package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256
// hashes. (See PopCount from Section 2.6.2.)

func main() {
	if len(os.Args) != 3 {
		return
	}
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%s / %s = %d\n", os.Args[1], os.Args[2], hashDifference(c1, c2))
	fmt.Printf("%s = %s\n", os.Args[1], c1)
	fmt.Printf("%s = %s\n", os.Args[2], c2)
}

func hashDifference(hash1 [32]byte, hash2 [32]byte) int {
	total := 0
	for i := 0; i < 32; i++ {
		byte1 := hash1[i]
		byte2 := hash2[i]
		for j := 0; j < 8; j++ {
			if byte1&1 != byte2&1 {
				total += 1
			}
		}
	}
	return total
}
