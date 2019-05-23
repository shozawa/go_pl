package evaluator

import (
	"testing"

	"github.com/shozawa/go_pl/ch08/lexer"
)

func TestParseLitral(t *testing.T) {
	prog := parse("42")
	if got := len(prog.statements); got != 1 {
		t.Errorf("len(prog.statements) is not 1. got=%v", got)
	}
	testFloatLiteral(t, prog.statements[0], 42)
}

func TestParseBinary(t *testing.T) {
	prog := parse("1 + 2")
	if got := len(prog.statements); got != 1 {
		t.Errorf("len(prog.statements) is not 1. got=%v", got)
	}
	b, ok := prog.statements[0].(binary)
	if !ok {
		t.Errorf("prog.statements[0] is not binary. got=%T", prog.statements[0])
	}
	if b.op != '+' {
		t.Errorf("bianry.op is not '+'. got=%v", b.op)
	}
	testFloatLiteral(t, b.x, 1)
	testFloatLiteral(t, b.y, 2)
}

func parse(input string) program {
	l := lexer.New(input)
	p := NewParser(l)
	return p.Parse()
}

func testFloatLiteral(t *testing.T, expr Expr, want float64) bool {
	e, ok := expr.(literal)
	if !ok {
		t.Errorf("expr is not literal. got=%T", e)
		return false
	}
	if float64(e) != want {
		t.Errorf("lit is not %v. got=%v", want, e)
	}
	return true
}
