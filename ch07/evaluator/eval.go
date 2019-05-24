package evaluator

import (
	"fmt"
	"math"

	"github.com/shozawa/go_pl/ch07/lexer"
)

func Eval(input string, env Env) float64 {
	l := lexer.New(input)
	p := NewParser(l)
	program := p.Parse()
	return program.Eval(env)
}

func (p program) Eval(env Env) float64 {
	var result float64
	for _, stmt := range p.statements {
		result = stmt.Eval(env)
	}
	return result
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	case '=':
		idnt, ok := b.x.(Var)
		if !ok {
			// FXIME: もっといいやり方がありそう
			panic("expect: Val")
		}
		env[idnt] = b.y.Eval(env)
		return env[idnt]
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "min":
		left := c.args[0].Eval(env)
		right := c.args[1].Eval(env)
		if left < right {
			return left
		} else {
			return right
		}
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
