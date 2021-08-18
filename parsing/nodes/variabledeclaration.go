package nodes

import (
	"fmt"
	"strings"
)

type VariableDeclaration struct {
	NodeType
	NodePosition
	Identifier string
	Expression Node
}

func CreateVariableDeclaration(position NodePosition, identifier string, expression Node) *VariableDeclaration {
	return &VariableDeclaration{
		NodeType:     NodeTypeVariableDeclaration,
		NodePosition: position,
		Identifier:   identifier,
		Expression:   expression,
	}
}

func (node *VariableDeclaration) String() string {
	var builder strings.Builder

	builder.WriteString("let ")
	builder.WriteString(node.Identifier)

	if node.Expression != nil {
		builder.WriteString(" = ")
		builder.WriteString(node.Expression.String())
	}

	builder.WriteString("\n")

	return builder.String()
}

func (node *VariableDeclaration) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "variable declaration")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}
