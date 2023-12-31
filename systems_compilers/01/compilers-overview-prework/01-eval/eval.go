package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
func Evaluate(expr ast.Expr) (int, error) {
	var result int
	switch ex := expr.(type) {
	case *ast.BasicLit:
		result = ReturnLiteralValue(ex)
	case *ast.BinaryExpr:
		result = EvalBinaryExpression(ex)
	case *ast.ParenExpr:
		result = EvalParenExpression(ex)
	default:
		fmt.Println("no match :(")
	}

	return result, nil
}

func ReturnLiteralValue(lit *ast.BasicLit) int {
	intVal, _ := strconv.Atoi(lit.Value)
	return intVal
}

func EvalBinaryExpression(bin *ast.BinaryExpr) int {
	left, _ := Evaluate(bin.X)
	right, _ := Evaluate(bin.Y)

	var eval int

	switch bin.Op {
	case token.ADD:
		eval = left + right
	case token.SUB:
		eval = left - right
	case token.MUL:
		eval = left * right
	}

	return eval
}

func EvalParenExpression(paren *ast.ParenExpr) int {
	val, _ := Evaluate(paren.X)
	return val
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
