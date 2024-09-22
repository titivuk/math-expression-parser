package lexer

import "github.com/titivuk/math-expression-parser/token"

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	return lexer
}

type Lexer struct {
	input string
	pos   int
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	ch := l.curChar()

	// leaves l.pos on the last char of the token
	switch {
	case ch == '+':
		tok = token.Token{
			Token:   token.PLUS,
			Literal: "+",
		}
	case ch == '-':
		tok = token.Token{
			Token:   token.MINUS,
			Literal: "-",
		}
	case ch == '*':
		tok = token.Token{
			Token:   token.MUL,
			Literal: "*",
		}
	case ch == '/':
		tok = token.Token{
			Token:   token.DIV,
			Literal: "/",
		}
	case isDigit(ch):
		start := l.pos
		end := l.pos
		for end+1 < len(l.input) && isDigit(l.input[end+1]) {
			end++
			l.pos++
		}

		tok = token.Token{
			Token:   token.VALUE,
			Literal: l.input[start : end+1],
		}
	case ch == 0:
		tok = token.Token{
			Token:   token.EOF,
			Literal: "",
		}
	}

	// advances l.pos to the next token
	l.pos++

	return tok
}

func (l *Lexer) curChar() byte {
	if len(l.input) <= l.pos {
		// EOF
		return 0
	}

	return l.input[l.pos]
}

func (l *Lexer) skipWhitespace() {

	for l.curChar() == ' ' || l.curChar() == '\t' || l.curChar() == '\n' || l.curChar() == '\r' {
		l.pos++
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
