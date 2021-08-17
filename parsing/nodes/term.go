package nodes

import (
	"fmt"
	"strings"
)

type Term struct {
	NodeType
	NodePosition
	Operator AdditiveOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=AdditiveOperator
type AdditiveOperator int

const (
	AdditiveOperatorAddition AdditiveOperator = iota
	AdditiveOperatorSubtraction
)

func CreateTerm(position NodePosition, operator AdditiveOperator, left Node, right Node) *Term {
	return &Term{
		NodeType:     NodeTypeTerm,
		NodePosition: position,
		Operator:     operator,
		Left:         left,
		Right:        right,
	}
}

func (node *Term) String() string {
	var builder strings.Builder

	builder.WriteString(node.Left.String())

	builder.WriteByte(' ')

	switch node.Operator {
	case AdditiveOperatorAddition:
		builder.WriteString("+")
	case AdditiveOperatorSubtraction:
		builder.WriteString("-")
	}

	builder.WriteByte(' ')

	builder.WriteString(node.Right.String())

	return builder.String()
}

func (node *Term) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"left\"];\n\"%p\" -> \"%p\" [label=\"right\"];\n%s%s", node, node.Operator.String(), node, node.Left, node, node.Right, node.Left.DotString(), node.Right.DotString())
}
