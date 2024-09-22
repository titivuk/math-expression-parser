package parser

import (
	"fmt"
	"strconv"

	"github.com/titivuk/math-expression-parser/ast"
	"github.com/titivuk/math-expression-parser/lexer"
	"github.com/titivuk/math-expression-parser/token"
)

type (
	prefixParseFn func() ast.Node
	infixParseFn  func(left ast.Node) ast.Node
)

const (
	_ = iota
	LOWEST
	VALUE
	SUM
	MUL
)

var precedences map[token.TokenType]int = map[token.TokenType]int{
	token.VALUE: VALUE,
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.MUL:   MUL,
	token.DIV:   MUL,
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l, prefixParseFns: make(map[token.TokenType]prefixParseFn), infixParseFns: make(map[token.TokenType]infixParseFn)}

	p.prefixParseFns[token.VALUE] = p.parseValueExpression

	p.infixParseFns[token.PLUS] = p.parseInfixExpression
	p.infixParseFns[token.MINUS] = p.parseInfixExpression
	p.infixParseFns[token.MUL] = p.parseInfixExpression
	p.infixParseFns[token.DIV] = p.parseInfixExpression

	p.nextToken()
	p.nextToken()

	return p
}

type Parser struct {
	lexer *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) Parse() ast.Node {
	expression := p.parseExpression(LOWEST)
	fmt.Printf("result expression: %v\n", expression)

	return expression
}

func (p *Parser) parseExpression(precedence int) ast.Node {
	prefixFn, ok := p.prefixParseFns[p.curToken.Token]
	if !ok {
		// TODO: aggregate errors
		fmt.Printf("no prefixFn for token: %s", p.curToken.Token)
		return nil
	}

	expression := prefixFn()
	fmt.Printf("prefixOrLeft: %v\n", expression)

	for p.peekToken.Token != token.EOF && precedence < precedences[p.peekToken.Token] {
		infixFn, ok := p.infixParseFns[p.peekToken.Token]
		if !ok {
			// TODO: aggregate errors
			fmt.Printf("no infixFn for token: %s", p.peekToken.Token)
			return nil
		}
		p.nextToken()

		expression = infixFn(expression)
	}

	return expression
}

func (p *Parser) parseValueExpression() ast.Node {
	value, _ := strconv.ParseFloat(p.curToken.Literal, 64)

	node := ast.ValueNode{Value: float64(value)}
	fmt.Printf("parseValueExpression: %v\n", node)

	return node
}

func (p *Parser) parseInfixExpression(left ast.Node) ast.Node {
	precedence := precedences[p.curToken.Token]
	node := ast.InfixNode{Left: left, Operator: p.curToken}
	p.nextToken()
	node.Right = p.parseExpression(precedence)
	fmt.Printf("parseInfixExpression: %v\n", node)

	return node
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
