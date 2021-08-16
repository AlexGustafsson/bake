package nodes

import (
	"strings"
)

type BinaryExpression struct {
	NodeType
	NodePosition
	Operator BinaryOperator
	Left     Node // BinaryExpression or UnaryExpression
	Right    Node // BinaryExpression or UnaryExpression
}

type BinaryOperator int

const (
	BinaryOperatorOr BinaryOperator = iota
	BinaryOperatorAnd
	BinaryOperatorEquals
	BinaryOperatorNotEquals
	BinaryOperatorLessThan
	BinaryOperatorLessThanOrEqual
	BinaryOperatorGreaterThan
	BinaryOperatorGreaterThanOrEqual
	BinaryOperatorAddition
	BinaryOperatorSubtraction
	BinaryOperatorMultiplication
	BinaryOperatorDivision
)

func CreateBinaryExpression(position NodePosition, operator BinaryOperator, left Node, right Node) *BinaryExpression {
	return &BinaryExpression{
		NodeType:     NodeTypeBinaryExpression,
		NodePosition: position,
		Operator:     operator,
		Left:         left,
		Right:        right,
	}
}

func (node *BinaryExpression) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case BinaryOperatorOr:
		builder.WriteString("||")
	case BinaryOperatorAnd:
		builder.WriteString("&&")
	case BinaryOperatorEquals:
		builder.WriteString("==")
	case BinaryOperatorNotEquals:
		builder.WriteString("!=")
	case BinaryOperatorLessThan:
		builder.WriteRune('<')
	case BinaryOperatorLessThanOrEqual:
		builder.WriteString("<=")
	case BinaryOperatorGreaterThan:
		builder.WriteRune('>')
	case BinaryOperatorGreaterThanOrEqual:
		builder.WriteString(">=")
	case BinaryOperatorAddition:
		builder.WriteRune('+')
	case BinaryOperatorSubtraction:
		builder.WriteRune('-')
	case BinaryOperatorMultiplication:
		builder.WriteRune('*')
	case BinaryOperatorDivision:
		builder.WriteRune('/')
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}
