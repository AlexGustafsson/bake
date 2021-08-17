package nodes

import (
	"fmt"
	"strings"
)

type Comparison struct {
	NodeType
	NodePosition
	Operator ComparisonOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=ComparisonOperator
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

func (node *Comparison) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"left\"];\n\"%p\" -> \"%p\" [label=\"right\"];\n%s%s", node, node.Operator.String(), node, node.Left, node, node.Right, node.Left.DotString(), node.Right.DotString())
}
