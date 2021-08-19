package nodes

import (
	"fmt"
	"strings"
)

type Term struct {
	NodeType
	Range
	Operator AdditiveOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=AdditiveOperator
type AdditiveOperator int

const (
	AdditiveOperatorAddition AdditiveOperator = iota
	AdditiveOperatorSubtraction
)

func CreateTerm(r Range, operator AdditiveOperator, left Node, right Node) *Term {
	return &Term{
		NodeType: NodeTypeTerm,
		Range:    r,
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (node *Term) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case AdditiveOperatorAddition:
		builder.WriteString("+")
	case AdditiveOperatorSubtraction:
		builder.WriteString("-")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}

func (node *Term) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
	builder.WriteString(node.Left.DotString())
	builder.WriteString(node.Right.DotString())
	return builder.String()
}
