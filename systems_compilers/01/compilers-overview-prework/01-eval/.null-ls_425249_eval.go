package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
func Evaluate(expr ast.Expr) (int, error) {
	//read the expression
	//execute in order
	//return result
	// TODO

	fmt.Printf("EXPR", expr)
	return 0, nil
}

func main() {
	expr, err := parser.ParseExpr("1 + 2 - 3 * 4")
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	err = ast.Print(fset, expr)
	if err != nil {
		log.Fatal(err)
	}
}
