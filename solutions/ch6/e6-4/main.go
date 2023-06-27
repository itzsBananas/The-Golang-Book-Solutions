// Exercise 6.4: Add a method Elems that returns a slice containing the elements of the set, suit-
// able for iterating over with a range loop.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Elems()) // "{1 9 144}"
}

func (s *IntSet) Elems() []int {
	elem := make([]int, s.Len())
	e := 0

	for i, word := range s.words {
		for j := uint(0); j < 64 && word > 0; j++ {
			if word&(1<<j) != 0 {
				elem[e] = (64*i + int(j))
				e += 1
			}
		}
	}

	return elem
}

// return the number of elements
func (s *IntSet) Len() int {
	n := 0
	for _, word := range s.words {
		for word != 0 {
			word = word & (word - 1) // clear rightmost non-zero bit
			n++
		}
	}
	return n
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	duplicate_words := make([]uint64, s.Len())
	copy(duplicate_words, s.words)

	return &IntSet{duplicate_words}
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
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

//!-string
