package expr

import "language/token"

// approximate abstract class in Java
type Node interface{}
type Expr interface {
	Accept(visitor Visitor)
}

type Literal struct {
	Value any
}

func (lit *Literal) Accept(visitor Visitor) {
	visitor.VisitLiteral(lit)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (bin *Binary) Accept(visitor Visitor) {
	visitor.VisitBinary(bin)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (unary *Unary) Accept(visitor Visitor) {
	visitor.VisitUnary(unary)
}

type Grouping struct {
	Expr Expr
}

func (group *Grouping) Accept(visitor Visitor) {
	visitor.VisitGrouping(group)
}

type Visitor interface {
	VisitBinary(binary *Binary)
	VisitGrouping(grouping *Grouping)
	VisitLiteral(literal *Literal)
	VisitUnary(unary *Unary)
}
