package ast

import (
	"strings"
)

type VariableDeclaration struct {
	NodeType
	Range
	Identifier string
	Expression Node
}

func CreateVariableDeclaration(r Range, identifier string, expression Node) *VariableDeclaration {
	return &VariableDeclaration{
		NodeType:   NodeTypeVariableDeclaration,
		Range:      r,
		Identifier: identifier,
		Expression: expression,
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
