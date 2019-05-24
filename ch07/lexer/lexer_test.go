package lexer

import (
	"testing"

	"github.com/shozawa/go_pl/ch07/token"
)

func TestSimple(t *testing.T) {
	l := Lexer{}
	got := l.NextToken()
	if got.Type != token.ASTERISK {
		t.Error("")
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		input string
		want  []token.Token
	}{
		{
			"42",
			[]token.Token{
				token.Token{Type: token.FLOAT, Literal: "42"},
				token.Token{Type: token.EOF, Literal: ""},
			},
		},
		{
			"1 + 1",
			[]token.Token{
				token.Token{Type: token.FLOAT, Literal: "1"},
				token.Token{Type: token.PLUS, Literal: "+"},
				token.Token{Type: token.FLOAT, Literal: "1"},
				token.Token{Type: token.EOF, Literal: ""},
			},
		},
		{
			`
			x = 16 
			sqrt(x)
			`,
			[]token.Token{
				token.Token{Type: token.INDT, Literal: "x"},
				token.Token{Type: token.ASSIGN, Literal: "="},
				token.Token{Type: token.FLOAT, Literal: "16"},
				token.Token{Type: token.INDT, Literal: "sqrt"},
				token.Token{Type: token.LPAREN, Literal: "("},
				token.Token{Type: token.INDT, Literal: "x"},
				token.Token{Type: token.RPAREN, Literal: ")"},
				token.Token{Type: token.EOF, Literal: ""},
			},
		},
		{
			"pow(2, 4)",
			[]token.Token{
				token.Token{Type: token.INDT, Literal: "pow"},
				token.Token{Type: token.LPAREN, Literal: "("},
				token.Token{Type: token.FLOAT, Literal: "2"},
				token.Token{Type: token.COMMA, Literal: ","},
				token.Token{Type: token.FLOAT, Literal: "4"},
				token.Token{Type: token.RPAREN, Literal: ")"},
				token.Token{Type: token.EOF, Literal: ""},
			},
		},
	}
	for _, test := range tests {
		l := New(test.input)
		for _, tok := range test.want {
			got := l.NextToken()
			if got.Type != tok.Type {
				t.Errorf("token.Type is not %q. got=%q", tok.Type, got.Type)
			}
			if got.Literal != tok.Literal {
				t.Errorf("token.Literal is not %q. got=%q", tok.Literal, got.Literal)
			}
		}
	}
}
