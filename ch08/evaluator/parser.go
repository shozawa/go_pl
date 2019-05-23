package evaluator

import (
	"strconv"

	"github.com/shozawa/go_pl/ch08/lexer"
	"github.com/shozawa/go_pl/ch08/token"
)

type (
	infixParseFn func(Expr) Expr
)

const (
	_ = iota
	LOWEST
	SUM
	PRODUCT
	CALL
)

var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.ASTERISK: PRODUCT,
	token.SLASH:    PRODUCT,
	token.LPAREN:   CALL,
}

type Parser struct {
	l             *lexer.Lexer
	curToken      token.Token
	peekToken     token.Token
	infixParseFns map[token.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseBinary)
	p.registerInfix(token.MINUS, p.parseBinary)
	p.registerInfix(token.ASTERISK, p.parseBinary)
	p.registerInfix(token.SLASH, p.parseBinary)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

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
	var left Expr
	switch p.curToken.Type {
	case token.FLOAT:
		left = p.parseFloatLiteral()
	case token.INDT:
		left = p.parseIdent()
	default:
		panic("error")
	}

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			// FIXME: どんなときに nil になる？ 必要があればエラー処理
			return left
		}
		p.nextToken()
		left = infix(left)
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

func (p *Parser) parseIdent() Expr {
	return Var(p.curToken.Literal)
}

func (p *Parser) parseBinary(left Expr) Expr {
	// FIXME: op を string 型にして余計な型変換をなくす
	b := binary{op: rune(p.curToken.Literal[0]), x: left}
	precedence := p.curPrecedence()
	p.nextToken() // cosume op
	b.y = p.parseExpression(precedence)
	return b
}

func (p *Parser) parseCallExpression(left Expr) Expr {
	ident, ok := left.(Var)
	if !ok {
		// FIXME: 適切にエラー処理する
		panic("error")
	}
	exp := call{fn: string(ident)}
	exp.args = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []Expr {
	var args []Expr

	p.nextToken() // consume '('

	// TODO: 複数の引数に対応する
	arg := p.parseExpression(LOWEST)
	args = append(args, arg)

	for p.peekTokenIs(token.COMMA) {

		p.nextToken() // consume prev arg
		p.nextToken() // consume ','

		arg := p.parseExpression(LOWEST)
		args = append(args, arg)
	}

	p.expectPeek(token.RPAREN) // ')'

	return args
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

func (p *Parser) registerInfix(tokenType token.TokenType, f infixParseFn) {
	p.infixParseFns[tokenType] = f
}

func (p *Parser) peekTokenIs(expect token.TokenType) bool {
	return p.peekToken.Type == expect
}

func (p *Parser) expectPeek(expect token.TokenType) bool {
	if p.peekToken.Type == expect {
		p.nextToken()
		return true
	} else {
		return false
	}
}
