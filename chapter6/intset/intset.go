/*
 * func (s *IntSet) Len() int      // returns the number of elements in set
 * func (s *IntSet) Remove(x int)  // remove element x from set
 * func (s *IntSet) Clear()        // empty the set
 * func (s *IntSet) Copy() *IntSet // return the copy
 */

package intset

import (
	"bytes"
	"fmt"
)

// An Intset is a set of small non-negative integers.
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

// String returns the set as a string of the form "{1,2,3}"
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

//!+ exam6-1
// Len returns the number of elements in set
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove delete x from set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) && s.words[word]&(1<<bit) != 0 {
		s.words[word] &= ^(1 << bit)
	}
}

// Clear
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy returns the copy
func (s *IntSet) Copy() *IntSet {
	tmp := make([]uint64, len(s.words))
	copy(tmp, s.words)
	return &IntSet{words: tmp}
}

//!-exam6-1

//!+exam6-2
// AddAll
func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		s.Add(x)
	}
}

//!-exam6-2

//!+exam6-3
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// IntersectionWith returns the intersection of s and t.
func (s *IntSet) IntersectionWith(t *IntSet) *IntSet {
	res := &IntSet{}
	l := min(len(s.words), len(t.words))
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]&t.words[i])
	}
	return res
}

// DifSet returns the difference between s and t.
func (s *IntSet) DifSet(t *IntSet) *IntSet {
	res := &IntSet{}
	l := min(len(s.words), len(t.words))
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]|(s.words[i]&t.words[i]))
	}
	res.words = append(res.words, s.words[l:]...)
	return res
}

// SymDifSet returns the element that's in one but not in other.
// XOR
func (s *IntSet) SymDifSet(t *IntSet) *IntSet {
	res := &IntSet{}
	l := min(len(s.words), len(t.words))
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]^(t.words[i]))
	}
	res.words = append(res.words, s.words[l:]...)
	res.words = append(res.words, t.words[l:]...)
	return res
}

//!-exam6-3

//!+exam6-4
// Elems returns the elements of set
func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, 64*i+j)
			}
		}
	}
	return res
}

//!-exam6-4

//!+exam6-5
// 32 << (^uint(0) >> 63) equals to 64 when in x86_64
// but 32 in x86

//!-exam6-5
