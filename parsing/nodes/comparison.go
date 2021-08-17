package nodes

import (
	"strings"
)

type Comparison struct {
	NodeType
	NodePosition
	Operator ComparisonOperator
	Left     Node
	Right    Node
}

type ComparisonOperator int

const (
	ComparisonOperatorOr ComparisonOperator = iota
	ComparisonOperatorAnd
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
	case ComparisonOperatorOr:
		builder.WriteString("||")
	case ComparisonOperatorAnd:
		builder.WriteString("&&")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}
