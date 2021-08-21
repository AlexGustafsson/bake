package nodes

import (
	"fmt"
	"strings"
)

type IfStatement struct {
	NodeType
	Range
	Expression     Node
	PositiveBranch Node
	// May be nil
	NegativeBranch Node
}

func CreateIfStatement(r Range, expression Node, positiveBranch Node, negativeBranch Node) *IfStatement {
	return &IfStatement{
		NodeType:       NodeTypeIfStatement,
		Range:          r,
		Expression:     expression,
		PositiveBranch: positiveBranch,
		NegativeBranch: negativeBranch,
	}
}

func (node *IfStatement) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "if %s ", node.Expression.String())
	builder.WriteString(node.PositiveBranch.String())
	if node.NegativeBranch != nil {
		fmt.Fprintf(&builder, " else %s", node.NegativeBranch.String())
	}
	return builder.String()
}
