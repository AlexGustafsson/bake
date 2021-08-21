package nodes

import (
	"strings"
)

type Comparison struct {
	NodeType
	Range
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

func CreateComparison(r Range, operator ComparisonOperator, left Node, right Node) *Comparison {
	return &Comparison{
		NodeType: NodeTypeComparison,
		Range:    r,
		Operator: operator,
		Left:     left,
		Right:    right,
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
