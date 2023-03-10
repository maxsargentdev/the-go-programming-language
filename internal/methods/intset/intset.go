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
	Words []uint64 // export this in my example and just use uint64 for compatability with cobra
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
		if i < len(s.Words) {
			s.Words[i] |= tword
		} else {
			s.Words = append(s.Words, tword)
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

func (*IntSet) Len() int {
	return 0
}

func (*IntSet) Remove(x int) {
	return
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
