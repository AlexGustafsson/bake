package nodes

import "strings"

type UnaryExpression struct {
	NodeType
	NodePosition
	Operator   UnaryOperator
	Expression Node // PrimaryExpression
}

type UnaryOperator int

const (
	UnaryOperatorSubtraction UnaryOperator = iota
	UnaryOperatorNot
	UnaryOperatorSpread
)

func CreateUnaryExpression(position NodePosition, operator UnaryOperator, expression Node) *UnaryExpression {
	return &UnaryExpression{
		NodeType:     NodeTypeUnaryExpression,
		NodePosition: position,
		Operator:     operator,
		Expression:   expression,
	}
}

func (node *UnaryExpression) String() string {
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
