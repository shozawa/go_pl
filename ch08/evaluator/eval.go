package evaluator

import "github.com/shozawa/go_pl/ch08/lexer"

func Eval(input string) float64 {
	l := lexer.New(input)
	p := NewParser(l)
	program := p.Parse()
	env := make(Env)
	return program.Eval(env)
}
