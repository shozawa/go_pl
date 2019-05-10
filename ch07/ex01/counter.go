package counter

import (
	"bufio"
	"bytes"
)

type WordCount int

func (c *WordCount) Write(p []byte) (int, error) {
	in := bytes.NewReader(p)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCount int

func (c *LineCount) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}
