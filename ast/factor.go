package ast

import (
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
