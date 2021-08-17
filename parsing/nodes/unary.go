package nodes

import "strings"

type Unary struct {
	NodeType
	NodePosition
	Operator   UnaryOperator
	Expression Node
}

type UnaryOperator int

const (
	UnaryOperatorSubtraction UnaryOperator = iota
	UnaryOperatorNot
	UnaryOperatorSpread
)

func CreateUnary(position NodePosition, operator UnaryOperator, expression Node) *Unary {
	return &Unary{
		NodeType:     NodeTypeUnary,
		NodePosition: position,
		Operator:     operator,
		Expression:   expression,
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

	builder.WriteString(node.Expression.String())

	return builder.String()
}
