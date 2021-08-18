package nodes

import (
	"fmt"
	"strings"
)

type Primary struct {
	NodeType
	NodePosition
	Operand Node
}

type PrimaryOperator int

func CreatePrimary(position NodePosition, operator PrimaryOperator, operand Node) *Primary {
	return &Primary{
		NodeType:     NodeTypePrimary,
		NodePosition: position,
		Operand:      operand,
	}
}

func (node *Primary) String() string {
	return node.Operand.String()
}

func (node *Primary) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "primary")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
	builder.WriteString(node.Operand.DotString())
	return builder.String()
}
