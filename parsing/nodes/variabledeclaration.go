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
	return fmt.Sprintf("\"%p\" [label=\"declaration %s\"];\n\"%p\" -> \"%p\" [label=\"expression\"];\n%s", node, node.Identifier, node, node.Expression, node.Expression.DotString())
}
