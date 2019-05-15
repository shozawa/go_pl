package strings

import "io"

type StringReader struct {
	input    string
	position int
}

func (r *StringReader) Read(p []byte) (int, error) {
	if r.position >= len(r.input) {
		return 0, io.EOF
	}
	for i := 0; i < len(r.input); i++ {
		p[i] = r.input[i]
		r.position++
	}
	return len(r.input), nil
}

func NewReader(input string) *StringReader {
	return &StringReader{input: input}
}
