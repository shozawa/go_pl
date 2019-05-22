package evaluator

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr Expr
		env  Env
		want float64
	}{
		{literal(5), nil, 5},
		{binary{op: '+', x: literal(5), y: literal(3)}, nil, 8},
		{binary{op: '+', x: literal(5), y: Var("x")}, Env{"x": 3}, 8},
	}
	for _, test := range tests {
		if got := test.expr.Eval(test.env); got != test.want {
			t.Errorf("want=%v, got=%v", test.want, got)
		}
	}
}
