package nodes

import (
	"strings"
)

type Equality struct {
	NodeType
	Range
	Operator EqualityOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=EqualityOperator
type EqualityOperator int

const (
	EqualityOperatorOr EqualityOperator = iota
	EqualityOperatorAnd
)

func CreateEquality(r Range, operator EqualityOperator, left Node, right Node) *Equality {
	return &Equality{
		NodeType: NodeTypeEquality,
		Range:    r,
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (node *Equality) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case EqualityOperatorOr:
		builder.WriteString("||")
	case EqualityOperatorAnd:
		builder.WriteString("&&")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}
