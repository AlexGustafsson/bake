package nodes

import (
	"fmt"
	"strings"
)

type Unary struct {
	NodeType
	NodePosition
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

func CreateUnary(position NodePosition, operator UnaryOperator, primary Node) *Unary {
	return &Unary{
		NodeType:     NodeTypeUnary,
		NodePosition: position,
		Operator:     operator,
		Primary:      primary,
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
	return fmt.Sprintf("\"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"Primary\"];\n%s", node, node.Operator.String(), node, node.Primary, node.Primary.DotString())
}
