package parser

import (
	"fmt"
	"language/expr"
	"language/token"
)

func HelloParser() {
	fmt.Println("Hello from the parser")
}

type Parser struct {
	Tokens  []token.Token
	Current int
}

func NewParser() Parser {
	return Parser{
		Tokens:  make([]token.Token, 0),
		Current: 0,
	}
}

// Private Methods
func (p *Parser) expression() expr.Expr {
	return p.equality()
}

func (p *Parser) equality() expr.Expr {
	expression := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expression = expr.Binary{
			Left:     expression,
			Operator: operator,
			Right:    right,
		}
	}

	return expr
}

func (p *Parser) comparison() expr.Expr {
	expr := p.term()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = expr.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

// HELPER FUNCTIONS
func (p *Parser) match(types ...token.TokenEnum) bool {
	for _, tok := range types {
		if p.check(tok) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tok token.TokenEnum) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tok
}

func (p *Parser) advance() {
	if !p.isAtEnd() {
		p.Current += 1
	}
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}
