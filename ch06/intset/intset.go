package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (i *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= 1 << bit
}

func (i *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(i.words) && i.words[word]&(1<<bit) != 0
}

func (i *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range i.words {
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
