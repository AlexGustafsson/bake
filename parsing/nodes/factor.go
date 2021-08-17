package nodes

import (
	"fmt"
	"strings"
)

type Factor struct {
	NodeType
	NodePosition
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

func CreateFactor(position NodePosition, operator MultiplicativeOperator, left Node, right Node) *Factor {
	return &Factor{
		NodeType:     NodeTypeFactor,
		NodePosition: position,
		Operator:     operator,
		Left:         left,
		Right:        right,
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
	return fmt.Sprintf("\"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"left\"];\n\"%p\" -> \"%p\" [label=\"right\"];\n%s%s", node, node.Operator.String(), node, node.Left, node, node.Right, node.Left.DotString(), node.Right.DotString())
}
