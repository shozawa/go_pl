package intset

import (
	"bytes"
	"fmt"
)

const (
	SIZE = 32 << (^uint(0) >> 63)
)

type IntSet struct {
	words []uint
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

func (s *IntSet) Elems() []uint {
	// FIXME: なんかもっと賢いやり方がありそう
	numbers := []uint{}
	for i := range s.words {
		for j := 0; j < SIZE; j++ {
			number := i*SIZE + j
			if s.Has(number) {
				numbers = append(numbers, uint(i*SIZE+j))
			}
		}
	}
	return numbers
}

func (s *IntSet) Add(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
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
	word, bit := x/SIZE, uint(x%SIZE)
	if word >= len(s.words) {
		// 何もしない
	} else {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/SIZE, uint(x%SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) UnionWith(other *IntSet) *IntSet {
	result := s.Copy()
	for i, word := range other.words {
		if i < len(result.words) {
			result.words[i] |= word
		} else {
			result.words = append(result.words, word)
		}
	}
	return result
}

func (s *IntSet) IntersectWith(other *IntSet) *IntSet {
	result := &IntSet{}
	// FIXME: もっときれいに
	for i := range s.words {
		if len(other.words) <= i {
			break
		}
		word := s.words[i] & other.words[i]
		result.words = append(result.words, word)
	}
	return result
}

func (s *IntSet) DifferenceWith(other *IntSet) *IntSet {
	result := s.Copy()
	for i := 0; i < min(len(s.words), len(other.words)); i++ {
		result.words[i] &= ^other.words[i]
	}
	return result
}

func (s *IntSet) SymetricDifferenceWith(other *IntSet) *IntSet {
	union := s.UnionWith(other)
	intersection := s.IntersectWith(other)
	return union.DifferenceWith(intersection)
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
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
		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", SIZE*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

func popcount(x uint) (count int) {
	for ; x != 0; x &= x - 1 {
		count++
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
