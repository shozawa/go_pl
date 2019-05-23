package evaluator

import (
	"strconv"

	"github.com/shozawa/go_pl/ch08/lexer"
	"github.com/shozawa/go_pl/ch08/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() program {
	prog := program{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseExpression()
		prog.statements = append(prog.statements, stmt)
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseExpression() Expr {
	switch p.curToken.Type {
	case token.FLOAT:
		value, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			// TODO: エラー処理
		}
		return literal(value)
	}
	panic("parse error")
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
