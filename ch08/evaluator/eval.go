package evaluator

import "github.com/shozawa/go_pl/ch08/lexer"

func Eval(input string, env Env) float64 {
	l := lexer.New(input)
	p := NewParser(l)
	program := p.Parse()
	return program.Eval(env)
}
