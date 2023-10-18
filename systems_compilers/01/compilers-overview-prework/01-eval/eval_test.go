package main

import (
	"go/parser"
	"testing"
)

func TestEvaluate(t *testing.T) {
	for _, test := range []struct {
		s        string
		expected int
	}{
		{"5", 5},
		{"2 * 3", 6},
		{"1 + 2 - 3 * 4", -9},

		{"2 * (3 - 4 * (5 + 6) + 7)", -68},
	} {
		//pass the string into the go parser
		//are returned an expr
		expr, err := parser.ParseExpr(test.s)
		if err != nil {
			t.Fatal(err)
		}
		//pass the resulting expr into our Evaluate function
		result, err := Evaluate(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result != test.expected {
			t.Fatalf("Expected %q to evaluate to %d, got %d", test.s, test.expected, result)
		}
	}
}
