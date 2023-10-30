package expr

import "language/token"

// approximate abstract class in Java
type Expr interface {
	Accept(Visitor) interface{}
}

type Literal struct {
	Value any
}

func (lit *Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteral(lit)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (bin *Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinary(bin)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (unary *Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnary(unary)
}

type Grouping struct {
	Expression Expr
}

func (group *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGrouping(group)
}

type Visitor interface {
	VisitBinary(binary *Binary) interface{}
	VisitGrouping(grouping *Grouping) interface{}
	VisitLiteral(literal *Literal) interface{}
	VisitUnary(unary *Unary) interface{}
}
