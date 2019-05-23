package evaluator

import (
	"strconv"

	"github.com/shozawa/go_pl/ch08/lexer"
	"github.com/shozawa/go_pl/ch08/token"
)

const (
	_ = iota
	LOWEST
	SUM
	PRODUCT
)

var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.ASTERISK: PRODUCT,
	token.SLASH:    PRODUCT,
}

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
		stmt := p.parseExpression(LOWEST)
		prog.statements = append(prog.statements, stmt)
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseExpression(precedence int) Expr {
	left := p.parseFloatLiteral()

	for precedence < p.peekPrecedence() {
		p.nextToken()
		left = p.parseBinary(left)
	}

	return left
}

func (p *Parser) parseFloatLiteral() Expr {
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		// TODO: エラー処理
	}
	return literal(value)
}

func (p *Parser) parseBinary(left Expr) Expr {
	// FIXME: op を string 型にして余計な型変換をなくす
	b := binary{op: rune(p.curToken.Literal[0]), x: left}
	precedence := p.curPrecedence()
	p.nextToken() // cosume op
	b.y = p.parseExpression(precedence)
	return b
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curPrecedence() int {
	return precedences[p.curToken.Type]
}

func (p *Parser) peekPrecedence() int {
	return precedences[p.peekToken.Type]
}
