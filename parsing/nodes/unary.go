package nodes

import (
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
