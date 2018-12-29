package main

import (
	"bytes"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
func test() {
	var x IntSet
	const cst = 1 << 32
	x.words = make([]uint64, 1<<27)
	// fmt.Println(len(x.words))
	// for i := 1; i < cst; i++ {
	// 	x.Add(i)
	// }
	// fmt.Println(count)
}

func main() {
	// defer trace("main")()
	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	test()
	pprof.StopCPUProfile()
	var x IntSet
	for i := 1; i < 11000; i++ {
		x.Add(i)
	}
	var y IntSet
	for i := 1; i < 1000; i++ {
		y.Add(i)
	}
	fmt.Println(x.words)
	fmt.Println(y.words)

	// x.Add(1)
	// x.Add(10000)
	// x.Remove(8)
	// fmt.Println(len(x.words))
	// fmt.Println(x.Len())
	// fmt.Println(y)
	// fmt.Println(x.words)
	// fmt.Println(x.String())
	// r := x.IntersectionWith(&y)
	r := x.SymDifSet(&y)
	fmt.Println(r.String())
	// fmt.Println(32 << (^uint(0) >> 63))
}

// An Intset is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

var count int

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
	s.words = []uint64{}
}

// Copy returns the copy
func (s *IntSet) Copy() *IntSet {
	tmp := make([]uint64, len(s.words))
	copy(tmp, s.words)
	return &IntSet{words: tmp}
}

// IntersectionWith returns the intersection of s and t.
func (s *IntSet) IntersectionWith(t *IntSet) *IntSet {
	res := &IntSet{}
	var l int
	if len(s.words) < len(t.words) {
		l = len(s.words)
	} else {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]&t.words[i])
	}
	return res
}

// SymDifSet returns the element that's in one but not in other.
func (s *IntSet) SymDifSet(t *IntSet) *IntSet {
	res := &IntSet{}
	l := min(len(s.words), len(t.words))
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]^(t.words[i]))
		// res.words = append(res.words, t.words[i]&(^(s.words[i])))
	}
	res.words = append(res.words, s.words[l:]...)
	res.words = append(res.words, t.words[l:]...)
	return res
}

func (s *IntSet) DifSet(t *IntSet) *IntSet {
	res := &IntSet{}
	l := min(len(s.words), len(t.words))
	for i := 0; i < l; i++ {
		res.words = append(res.words, s.words[i]|(s.words[i]&t.words[i]))
	}
	// res.words = append(res.words, s.words[l:]...)
	// res.words = append(res.words, t.words[l:]...)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
