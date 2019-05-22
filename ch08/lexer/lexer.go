package lexer

import (
	"github.com/shozawa/go_pl/ch08/token"
)

func New(input string) *Lexer {
	return &Lexer{input: input}
}

type Lexer struct {
	input    string
	position int
	ch       byte
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{Type: token.ASTERISK, Literal: "*"}
}
