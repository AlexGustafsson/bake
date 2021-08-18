package nodes

import (
	"fmt"
	"strings"
)

type Comparison struct {
	NodeType
	NodePosition
	Operator ComparisonOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=ComparisonOperator
type ComparisonOperator int

const (
	ComparisonOperatorEquals ComparisonOperator = iota
	ComparisonOperatorNotEquals
	ComparisonOperatorLessThan
	ComparisonOperatorLessThanOrEqual
	ComparisonOperatorGreaterThan
	ComparisonOperatorGreaterThanOrEqual
)

func CreateComparison(position NodePosition, operator ComparisonOperator, left Node, right Node) *Comparison {
	return &Comparison{
		NodeType:     NodeTypeComparison,
		NodePosition: position,
		Operator:     operator,
		Left:         left,
		Right:        right,
	}
}

func (node *Comparison) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case ComparisonOperatorEquals:
		builder.WriteString("==")
	case ComparisonOperatorNotEquals:
		builder.WriteString("!=")
	case ComparisonOperatorLessThan:
		builder.WriteRune('<')
	case ComparisonOperatorLessThanOrEqual:
		builder.WriteString("<=")
	case ComparisonOperatorGreaterThan:
		builder.WriteRune('>')
	case ComparisonOperatorGreaterThanOrEqual:
		builder.WriteString(">=")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}

func (node *Comparison) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
	builder.WriteString(node.Left.DotString())
	builder.WriteString(node.Right.DotString())
	return builder.String()
}
