package nodes

import (
	"fmt"
	"strings"
)

type Factor struct {
	NodeType
	Range
	Operator MultiplicativeOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=MultiplicativeOperator
type MultiplicativeOperator int

const (
	MultiplicativeOperatorMultiplication MultiplicativeOperator = iota
	MultiplicativeOperatorDivision
)

func CreateFactor(r Range, operator MultiplicativeOperator, left Node, right Node) *Factor {
	return &Factor{
		NodeType: NodeTypeFactor,
		Range:    r,
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (node *Factor) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case MultiplicativeOperatorMultiplication:
		builder.WriteString("*")
	case MultiplicativeOperatorDivision:
		builder.WriteString("/")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}

func (node *Factor) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
	builder.WriteString(node.Left.DotString())
	builder.WriteString(node.Right.DotString())
	return builder.String()
}
