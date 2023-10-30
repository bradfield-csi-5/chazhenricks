package tools

import (
	"fmt"
	"language/expr"
	"strings"
)

type astPrinter struct{}

func PrintAst(expression expr.Expr) string {
	printer := astPrinter{}
	return printer.Print(expression)
}

func (printer astPrinter) Print(expression expr.Expr) string {
	return expression.Accept(printer).(string)
}

func (printer astPrinter) VisitBinary(binary *expr.Binary) interface{} {
	return printer.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}
func (printer astPrinter) VisitGrouping(grouping *expr.Grouping) interface{} {

	return printer.parenthesize("group", grouping.Expression)
}

func (printer astPrinter) VisitLiteral(literal *expr.Literal) interface{} {
	return fmt.Sprintf("%v", literal.Value)
}

func (printer astPrinter) VisitUnary(unary *expr.Unary) interface{} {
	return printer.parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (printer astPrinter) parenthesize(name string, parts ...interface{}) string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(name)

	for _, part := range parts {
		str.WriteString(" ")
		switch p := part.(type) {
		case expr.Expr:
			str.WriteString(p.Accept(printer).(string))
		case string:
			str.WriteString(p)
		case fmt.Stringer:
			str.WriteString(p.String())
		}
	}
	str.WriteString(")")

	return str.String()
}
