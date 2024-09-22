package evaluator

import (
	"fmt"

	"github.com/titivuk/math-expression-parser/ast"
	"github.com/titivuk/math-expression-parser/object"
	"github.com/titivuk/math-expression-parser/token"
)

func Eval(node ast.Node) object.Object {
	switch n := node.(type) {
	case ast.ValueNode:
		return &object.Result{Value: n.Value}
	case ast.InfixNode:
		left := Eval(n.Left)
		if isError(left) {
			return left
		}

		right := Eval(n.Right)
		if isError(right) {
			return right
		}

		return evalInfixExpression(left, right, n.Operator)
	}

	return nil
}

func evalInfixExpression(left, right object.Object, operator token.Token) object.Object {
	if left.Type() != right.Type() {
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	leftValue := left.(*object.Result).Value
	rightValue := right.(*object.Result).Value

	switch operator.Token {
	case token.PLUS:
		return &object.Result{Value: leftValue + rightValue}
	case token.MINUS:
		return &object.Result{Value: leftValue - rightValue}
	case token.MUL:
		return &object.Result{Value: leftValue * rightValue}
	case token.DIV:
		return &object.Result{Value: leftValue / rightValue}
	}

	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func isError(obj object.Object) bool {
	return obj.Type() == object.ERROR_OBJ
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
