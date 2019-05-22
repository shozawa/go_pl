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
	EOF      = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}
