package lexer

import (
	"testing"

	"github.com/shozawa/go_pl/ch08/token"
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
			"1 + 1",
			[]token.Token{
				token.Token{Type: token.FLOAT, Literal: "1"},
				token.Token{Type: token.PLUS, Literal: "+"},
				token.Token{Type: token.FLOAT, Literal: "1"},
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
