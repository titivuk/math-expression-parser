package token

const (
	VALUE = "VALUE"

	PLUS  = "+"
	MINUS = "-"
	MUL   = "*"
	DIV   = "/"

	LPAREN = "("
	RPAREN = ")"

	EOF = "EOF"
)

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}
