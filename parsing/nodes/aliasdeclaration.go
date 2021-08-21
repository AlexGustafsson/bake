package nodes

import (
	"fmt"
	"strings"
)

type AliasDeclaration struct {
	NodeType
	Range
	Identifier string
	Expression Node
}

func CreateAliasDeclaration(r Range, identifier string, expression Node) *AliasDeclaration {
	return &AliasDeclaration{
		NodeType:   NodeTypeAliasDeclaration,
		Range:      r,
		Identifier: identifier,
		Expression: expression,
	}
}

func (node *AliasDeclaration) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "alias %s : ", node.Identifier)
	builder.WriteString(node.Expression.String())
	builder.WriteRune('\n')

	return builder.String()
}
