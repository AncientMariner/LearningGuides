package behavioral

import (
	"fmt"
	"strconv"
	"strings"
)

// Interpreter: defines a grammar and interprets sentences in it.
// Rarely implemented from scratch in Go — prefer parser libs (participle, antlr).
// This example interprets simple arithmetic expressions: "3 + 5 - 2"

type Expression interface {
	Interpret() int
}

type NumberExpr struct{ value int }

func (n *NumberExpr) Interpret() int { return n.value }

type AddExpr struct{ left, right Expression }

func (a *AddExpr) Interpret() int { return a.left.Interpret() + a.right.Interpret() }

type SubExpr struct{ left, right Expression }

func (s *SubExpr) Interpret() int { return s.left.Interpret() - s.right.Interpret() }

// Parse parses "num op num op num ..." left to right.
func Parse(input string) Expression {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return &NumberExpr{}
	}
	n, _ := strconv.Atoi(tokens[0])
	var expr Expression = &NumberExpr{value: n}
	for i := 1; i+1 < len(tokens); i += 2 {
		right, _ := strconv.Atoi(tokens[i+1])
		rightExpr := &NumberExpr{value: right}
		switch tokens[i] {
		case "+":
			expr = &AddExpr{left: expr, right: rightExpr}
		case "-":
			expr = &SubExpr{left: expr, right: rightExpr}
		default:
			fmt.Printf("unknown operator: %s\n", tokens[i])
		}
	}
	return expr
}

func ExampleInterpreter() {
	expr := Parse("3 + 5 - 2")
	fmt.Println(expr.Interpret())
	// Output:
	// 6
}
