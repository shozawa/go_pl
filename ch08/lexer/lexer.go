package lexer

import (
	"github.com/shozawa/go_pl/ch08/token"
)

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ch = l.input[0]
	return l
}

type Lexer struct {
	input    string
	position int
	ch       byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch {
	case isDigit(l.ch):
		return l.readDigit()
	case isIdentifier(l.ch):
		return l.readIdentifier()
	}

	switch l.ch {
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: "-"}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: "*"}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: "/"}
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case 0:
		tok = token.Token{Type: token.EOF}
	default:
		tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readDigit() token.Token {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	literal := l.input[start:l.position]
	return token.Token{Type: token.FLOAT, Literal: literal}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.position
	for isIdentifier(l.ch) {
		l.readChar()
	}
	literal := l.input[start:l.position]
	return token.Token{Type: token.INDT, Literal: literal}
}

func (l *Lexer) readChar() {
	l.position += 1
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isIdentifier(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
