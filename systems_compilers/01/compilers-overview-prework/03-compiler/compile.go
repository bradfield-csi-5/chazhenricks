package main

import (
	"fmt"
	"go/ast"
	"go/token"
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

	// fmt.Printf("NODE: %v\n", node)
	//handle arguments
	bodyStatements := node.Body.List
	for _, stmt := range bodyStatements {
		EvaluateStatement(stmt)
	}

	//pop the "return" memory address and add a halt to stop the program
	builder.WriteString("pop 0\n")
	builder.WriteString("halt\n")
	result := builder.String()
	// fmt.Printf("IM THE ASSEMBLY:%s\n", result)

	return result, nil
}

func EvaluateStatement(stmt ast.Stmt) {
	switch stmtType := stmt.(type) {
	case *ast.ReturnStmt:
		EvaluateReturnStatement(stmtType)
	}
}

func EvaluateExpression(expr ast.Expr) {
	switch ex := expr.(type) {
	case *ast.BasicLit:
		EvalLiteralExpr(ex)
	case *ast.BinaryExpr:
		EvalBinaryExpr(ex)
	case *ast.Ident: //when we just get x or y
		EvalIdent(ex)
	case *ast.ParenExpr:
		EvalParensExpr(ex)

	default:
		fmt.Printf("WHAT AM I %T\n", ex)
	}

}

func EvalParensExpr(paren *ast.ParenExpr) {
	EvaluateExpression(paren.X)
}

func EvaluateReturnStatement(ret *ast.ReturnStmt) {

	switch retExpr := ret.Results[0].(type) {
	case *ast.BasicLit:
		EvalLiteralExpr(retExpr)
	case *ast.BinaryExpr:
		EvalBinaryExpr(retExpr)
	}

}

func EvalLiteralExpr(expr *ast.BasicLit) {
	PushValueToStack(expr.Value)
}

func GetLiteralValue(expr *ast.BasicLit) string {
	return expr.Value
}

func EvalBinaryExpr(expr *ast.BinaryExpr) {
	EvaluateExpression(expr.X)
	EvaluateExpression(expr.Y)
	EvalOperand(expr.Op)
}

func EvalOperand(op token.Token) {
	switch op {
	case token.ADD:
		builder.WriteString("add\n")
	case token.SUB:
		builder.WriteString("sub\n")
	case token.MUL:
		builder.WriteString("mul\n")
	case token.QUO:
		builder.WriteString("div\n")
	}
}

func EvalIdent(expr *ast.Ident) {
	fmt.Printf("NAME: %s\n", expr.Name)
	switch expr.Name {
	case "x":
		//pop from memory address 1 and add to stack
		builder.WriteString("push 1\n")
	case "y":
		//pop from memory address 2 and add to stack
		builder.WriteString("push 2\n")
	default:
		fmt.Println("You have fucked up with the variables")
	}
}

func PushValueToStack(value string) {
	builder.WriteString("pushi ")
	builder.WriteString(value)
	builder.WriteString("\n")
}
