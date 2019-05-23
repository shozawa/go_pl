package token

type TokenType string

const (
	INDT     = "INDT"
	FLOAT    = "FLOAT"
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LPAREN   = "("
	RPAREN   = ")"
	COMMA    = ","
	EOF      = "EOF"
	ILLEGAL  = "ILLEGAL"
)

type Token struct {
	Type    TokenType
	Literal string
}
