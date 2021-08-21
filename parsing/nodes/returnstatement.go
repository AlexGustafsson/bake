package nodes

import (
	"fmt"
	"strings"
)

type ReturnStatement struct {
	NodeType
	Range
	Value Node
}

func CreateReturnStatement(r Range, value Node) *ReturnStatement {
	return &ReturnStatement{
		NodeType: NodeTypeReturnStatement,
		Range:    r,
		Value:    value,
	}
}

func (node *ReturnStatement) String() string {
	return fmt.Sprintf("return %s\n", node.Value.String())
}

func (node *ReturnStatement) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"return\"];\n", node)
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"value\"];\n", node, node.Value)
	builder.WriteString(node.Value.DotString())
	return builder.String()
}
