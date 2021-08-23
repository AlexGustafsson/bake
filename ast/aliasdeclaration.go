package ast

import (
	"fmt"
	"strings"
)

type AliasDeclaration struct {
	NodeType
	Range
	Exported   bool
	Identifier string
	Expression Node
}

func CreateAliasDeclaration(r Range, exported bool, identifier string, expression Node) *AliasDeclaration {
	return &AliasDeclaration{
		NodeType:   NodeTypeAliasDeclaration,
		Range:      r,
		Exported:   exported,
		Identifier: identifier,
		Expression: expression,
	}
}

func (node *AliasDeclaration) String() string {
	var builder strings.Builder

	if node.Exported {
		builder.WriteString("export ")
	}

	fmt.Fprintf(&builder, "alias %s : ", node.Identifier)
	builder.WriteString(node.Expression.String())
	builder.WriteRune('\n')

	return builder.String()
}
