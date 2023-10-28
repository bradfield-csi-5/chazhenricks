package expr

import "language/token"

// approximate abstract class in Java
type Node interface{}
type Expr interface {
	Node
}

type Literal struct {
	Kind token.Token
}
