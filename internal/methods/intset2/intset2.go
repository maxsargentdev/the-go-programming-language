package intset2

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	Words []uint // export this in my example
}

const uintSize = 32 << (^uint(0) >> 63)

// 4bytes = uint32
// 8bytes = uint64
//
//^uint(0) is the uint value in which all bits are set.
//Right-shifting the result of the first step by 63 yields
//
// 1 bit set on a 64 bit platform
// 0 bits set on a 32 bit platform
//
//
//Right shift all the way 63, a 32 bit word of 1s will clear and a 64bit word of 1s will be 1 bit left (the bit representing 32 prev)
//Left shift back 32 will create a value of either 32 or 64
//
//0 on a 32-bit architecture, and
//1 on a 64-bit architecture.
//
//Left-shifting 32 by as many places as the result of the second step yields
//
//32 on a 32-bit architecture, and
//64 on a 64-bit architecture.
//
//We are detecting whether or not its possible for us to set the 33rd bit, which has a denary value of 64

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
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
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
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
func popcount(x uint) (count int) {
	for x != 0 {
		count++
		x &= x - 1
	}
	return
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/uintSize, uint(x%uintSize)
		s.Words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.Words = nil
}

func (s *IntSet) Copy() *IntSet {
	var copyOfWords []uint

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

func (s *IntSet) Elems() []int {
	if s.Len() == 0 {
		return nil // set is empty and this is the zero value of a slice
	}
	elements := make([]int, 0, s.Len())

	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, uintSize*i+j)
			}
		}
	}

	return elements

}
