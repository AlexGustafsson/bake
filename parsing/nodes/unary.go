package nodes

import (
	"fmt"
	"strings"
)

type Unary struct {
	NodeType
	Range
	Operator UnaryOperator
	Primary  Node
}

//go:generate stringer -type=UnaryOperator
type UnaryOperator int

const (
	UnaryOperatorSubtraction UnaryOperator = iota
	UnaryOperatorNot
	UnaryOperatorSpread
)

func CreateUnary(r Range, operator UnaryOperator, primary Node) *Unary {
	return &Unary{
		NodeType: NodeTypeUnary,
		Range:    r,
		Operator: operator,
		Primary:  primary,
	}
}

func (node *Unary) String() string {
	var builder strings.Builder

	switch node.Operator {
	case UnaryOperatorSubtraction:
		builder.WriteRune('-')
	case UnaryOperatorNot:
		builder.WriteRune('!')
	case UnaryOperatorSpread:
		builder.WriteString("...")
	}

	builder.WriteString(node.Primary.String())

	return builder.String()
}

func (node *Unary) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "unary")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Primary, "primary")
	builder.WriteString(node.Primary.DotString())
	return builder.String()
}
