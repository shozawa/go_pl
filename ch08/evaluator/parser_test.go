package evaluator

import (
	"testing"

	"github.com/shozawa/go_pl/ch08/lexer"
)

func TestParseLiteral(t *testing.T) {
	l := lexer.New("42")
	p := NewParser(l)
	prog := p.Parse()
	if got := len(prog.statements); got != 1 {
		t.Errorf("len(prog.statements) is not 1. got=%v", got)
	}
	lit, ok := prog.statements[0].(literal)
	if !ok {
		t.Errorf("prog.statements[0] is not literal. got=%T", prog.statements[0])
	}
	if lit != literal(42) {
		t.Errorf("lit is not literal(42). got=%v", lit)
	}
}
