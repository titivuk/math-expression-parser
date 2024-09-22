package lexer

import (
	"testing"

	"github.com/titivuk/math-expression-parser/token"
)

func TestNextToken(t *testing.T) {
	input := "5 + 12+3 - 2"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"VALUE", "5"},
		{"+", "+"},
		{"VALUE", "12"},
		{"+", "+"},
		{"VALUE", "3"},
		{"-", "-"},
		{"VALUE", "2"},
		{"EOF", ""},
	}

	lexer := NewLexer(input)

	for _, tt := range tests {
		tok := lexer.NextToken()

		if tok.Token != tt.expectedType {
			t.Fatalf("Invalid token type - \"%s\". Expected - \"%s\"", tok.Token, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Invalid token literal - \"%s\". Expected - \"%s\"", tok.Literal, tt.expectedLiteral)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input   byte
		isDigit bool
	}{
		{'0', true},
		{'9', true},
		{'+', false},
		{'-', false},
	}

	for _, tt := range tests {
		if isDigit(tt.input) != tt.isDigit {
			t.Fatalf("isDigit check failed for '%c'. Expected - %t. Actual - %t", tt.input, tt.isDigit, !tt.isDigit)
		}
	}
}
