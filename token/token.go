package token

const (
	VALUE = "VALUE"

	PLUS  = "+"
	MINUS = "-"
	MUL   = "*"
	DIV   = "/"

	EOF = "EOF"
)

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}
