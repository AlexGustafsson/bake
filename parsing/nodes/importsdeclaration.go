package nodes

import (
	"fmt"
	"strings"
)

type ImportsDeclaration struct {
	NodeType
	Range
	Imports []*InterpretedString
}

func CreateImportsDeclaration(r Range, imports []*InterpretedString) *ImportsDeclaration {
	return &ImportsDeclaration{
		NodeType: NodeTypeImportsDeclaration,
		Range:    r,
		Imports:  imports,
	}
}

func (node *ImportsDeclaration) String() string {
	var builder strings.Builder

	builder.WriteString("import (\n")

	for _, node := range node.Imports {
		builder.WriteString(node.String())
		builder.WriteRune('\n')
	}

	builder.WriteString(")\n")

	return builder.String()
}

func (node *ImportsDeclaration) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "imports")
	for _, literal := range node.Imports {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, literal)
		builder.WriteString(literal.DotString())
	}
	return builder.String()
}
