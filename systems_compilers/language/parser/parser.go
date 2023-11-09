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
		expression = &expr.Binary{
			Left:     expression,
			Operator: operator,
			Right:    right,
		}
	}

	return expression
}

func (p *Parser) comparison() expr.Expr {
	expression := p.term()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}
	return expression
}

func (p *Parser) term() expr.Expr {
	expression := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.unary()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}
	return expression
}

func (p *Parser) factor() expr.Expr {
	expression := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expression = &expr.Binary{Left: expression, Operator: operator, Right: right}
	}
	return expression
}

func (p *Parser) unary() expr.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &expr.Unary{Operator: operator, Right: right}
	}

	return p.primary()
}

func (p *Parser) primary() expr.Expr {
	if p.match(token.FALSE) {
		return &expr.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return &expr.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return &expr.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		return &expr.Literal{Value: p.previous().Literal}
	}
	if p.match(token.LEFT_PAREN) {
		expression := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ( after expression")
		return &expr.Grouping{Expression: expression}
	}

	//if none of these match, throw
  return 
}
// Error Handling

type ParseError struct {
	Token   token.TokenEnum
	Message string
}

func newParseError(token token.TokenEnum, message string) error {
	return &ParseError{Token: token, Message: message}
}



func (e *ParseError) Error() string {
	return e.Message
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

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current += 1
	}
	return p.previous()
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

func (p *Parser) consume(tok token.TokenEnum, message string) token.Token {

	if !p.check(tok) {
		//throw error
    err :=p.handleError(tok, message)
    panic(err)
	}

	return p.advance()
}

func (p *Parser) handleError(token token.TokenEnum, message string)error {
  return newParseError(token, message)
}
