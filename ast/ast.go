package ast

import "github.com/titivuk/math-expression-parser/token"

type Node interface {
}

type ValueNode struct {
	Value float64
}

type InfixNode struct {
	Operator token.Token
	Left     Node
	Right    Node
}
