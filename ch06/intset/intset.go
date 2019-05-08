package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func New(numbers ...int) *IntSet {
	set := &IntSet{}
	for _, n := range numbers {
		set.Add(n)
	}
	return set
}

func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		count += popcount(word)
	}
	return
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(numbers ...int) {
	for _, n := range numbers {
		s.Add(n)
	}	
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		// 何もしない
	} else {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) UnionWith(other *IntSet) {
	for i, word := range other.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	words := make([]uint64, len(s.words))
	copy(words, s.words)
	dst := &IntSet{}
	dst.words = words
	return dst
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

func popcount(x uint64) (count int) {
	for ; x != 0; x &= x - 1 {
		count++
	}
	return
}
