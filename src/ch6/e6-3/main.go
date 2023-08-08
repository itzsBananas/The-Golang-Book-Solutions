// Exercise 6.3: (*IntSet).UnionWith computes the union of two sets using |, the word-paral-
// lel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and Sym-
// metricDifference for the corresponding set operations. (The symmetric difference of two168
// sets contains the elements present in one set or the other but not both.)

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.SymmetricDifference(&y)
	fmt.Println(x.String())
}

func (s *IntSet) IntersectWith(t *IntSet) {
	length := len(s.words)
	if length > len(t.words) {
		length = len(t.words)
	}
	for i := 0; i < length; i++ {
		s.words[i] &= t.words[i]
	}
	for i := length; i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	length := len(s.words)
	if length > len(t.words) {
		length = len(t.words)
	}
	for i := 0; i < length; i++ {
		s.words[i] &= ^t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	length := len(s.words)
	if length > len(t.words) {
		length = len(t.words)
	}
	for i := 0; i < length; i++ {
		s.words[i] ^= t.words[i]
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

type IntSet struct {
	words []uint64
}
