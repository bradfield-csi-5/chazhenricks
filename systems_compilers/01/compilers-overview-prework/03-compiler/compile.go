package main

import (
	"fmt"
	"go/ast"
	"strings"
)

var builder strings.Builder

// Given an AST node corresponding to a function (guaranteed to be
// of the form `func f(x, y byte) byte`), compile it into assembly
// code.
//
// Recall from the README that the input parameters `x` and `y` should
// be read from memory addresses `1` and `2`, and the return value
// should be written to memory address `0`.
func compile(node *ast.FuncDecl) (string, error) {

	bodyStatements := node.Body.List
	for _, stmt := range bodyStatements {
		EvaluateStatement(stmt)
	}

	//pop the "return" memory address and add a halt to stop the program
	builder.WriteString("pop 0\n")
	builder.WriteString("halt\n")
	result := builder.String()

	return result, nil
}

func EvaluateStatement(stmt ast.Stmt) {
	switch stmtType := stmt.(type) {
	case *ast.ReturnStmt:
		getValueFromReturnStmt(stmtType)
	}
}

func EvaluateExpression(expr ast.Expr) {
	switch ex := expr.(type) {
	case *ast.BasicLit:
		EvalLiteralValue(ex)
	}
}

func getValueFromReturnStmt(ret *ast.ReturnStmt) {

	switch retExpr := ret.Results[0].(type) {
	case *ast.BasicLit:
		EvalLiteralValue(retExpr)
	case *ast.BinaryExpr:
		fmt.Println("IM A BINARY")
		// EvaluateBinaryExpression(retExpr)
	}

}

func EvalLiteralValue(expr *ast.BasicLit) {
	PushValueToStack(expr.Value)
}

func PushValueToStack(value string) {
	fmt.Println("I GOT PUSHED TO")
	builder.WriteString("pushi ")
	builder.WriteString(value)
	builder.WriteString("\n")
}
