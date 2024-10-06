package main

import (
	"github.com/titivuk/math-expression-parser/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
