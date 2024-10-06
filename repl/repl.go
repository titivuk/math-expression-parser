package repl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/titivuk/math-expression-parser/evaluator"
	"github.com/titivuk/math-expression-parser/lexer"
	"github.com/titivuk/math-expression-parser/object"
	"github.com/titivuk/math-expression-parser/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">> ")

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)
		tree := p.Parse()

		result := evaluator.Eval(tree)

		switch result := result.(type) {
		case *object.Error:
			io.WriteString(out, fmt.Sprintf("Error result: %s", result.Message))
		case *object.Result:
			io.WriteString(out, strconv.FormatFloat(result.Value, 'f', -1, 64))
		default:
			io.WriteString(out, fmt.Sprintf("Unknown error: %+v", result))
		}

		io.WriteString(out, "\n")
	}
}
