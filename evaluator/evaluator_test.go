package evaluator

import (
	"testing"

	"github.com/titivuk/math-expression-parser/lexer"
	"github.com/titivuk/math-expression-parser/object"
	"github.com/titivuk/math-expression-parser/parser"
)

func TestEval(t *testing.T) {
	tests := []struct {
		Input  string
		Output float64
	}{
		{Input: "1+2", Output: 3},
		{Input: "5-3", Output: 2},
		{Input: "0-54", Output: -54},
		{Input: "1+16/2-4+5*3", Output: 20},
		{Input: "(1+2)", Output: 3},
		{Input: "2*(1+3)", Output: 8},
		{Input: "(1+3)*2", Output: 8},
		{Input: "(1+2+3)/3", Output: 2},
		{Input: "((1+2)*4)/3", Output: 4},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.Input)
		p := parser.NewParser(l)
		tree := p.Parse()

		result := Eval(tree)

		switch result := result.(type) {
		case *object.Error:
			t.Errorf("Error result: %s", result.Message)
		case *object.Result:
			if result.Value != tt.Output {
				t.Errorf("Invalid result value. Expected: %f. Actual: %f", tt.Output, result.Value)
			}
		default:
			t.Errorf("Unknown error: %+v", result)
		}
	}
}
