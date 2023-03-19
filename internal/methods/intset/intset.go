// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	Words []uint64 // export this in my example
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) { // for all the words that also exist in s
			s.Words[i] &= tword // calculate the intersection with bitwise OR (AuB)
		} else { // for the words that dont exist in s
			s.Words = append(s.Words, tword) // slap em on the end, this is a UNION after all
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint64(j)) != 0 {
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

// Returns the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.Words {
		count += popcount(word)
	}
	return count
}

// Returns number of 1 bits
func popcount(x uint64) (count int) {
	for x != 0 {
		count++
		x &= x - 1
	}
	return
}

func (s *IntSet) Remove(x int) {
	//word, bit := x/64, uint(x%64)
	//fmt.Println("word:", word)
	//fmt.Println("bit (0 index):", bit)
	//fmt.Printf("integer: %d\n", uint64(s.Words[bit]))
	//
	//fmt.Printf("binary: %b\n", ^uint64(0))
	//fmt.Printf("binary: %b\n", (uint64(1) << bit))
	//fmt.Printf("binary: %b\n", (^uint64(0))^(uint64(1)<<bit))
	//fmt.Printf("binary: %b\n", s.Words)
	//mask := (^uint64(0)) ^ (uint64(1) << bit)
	//
	//s.Words[word] &= mask
	//fmt.Printf("%v\n", *s)
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.Words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.Words = nil
}

func (s *IntSet) Copy() *IntSet {
	var copyOfWords []uint64

	for _, ele := range s.Words {
		copyOfWords = append(copyOfWords, ele)
	}

	return &IntSet{Words: copyOfWords}
}

func (s *IntSet) AddAll(intss ...int) {
	for _, v := range intss {
		s.Add(v)
	}
}

// IntersectWith - Calculates the intersection of the two sets
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &= tword // AND the words (AnB)
		} else {
			s.Words = append(s.Words, tword) //
		}
	}
}

// DifferenceWith - Calculates the difference of the two sets
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &^= tword // AND NOT (bit clear) (A n B') What is in A AND NOT IN B
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}

// SymmetricDifference - Calculates the symettricdifference of the two sets
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] ^= tword // XOR (exclusive or) union of members which appear only in one of either set
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}
