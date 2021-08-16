package nodes

import (
	"strings"
)

type ImportsDeclaration struct {
	NodeType
	NodePosition
	Imports []*InterpretedString
}

func CreateImportsDeclaration(position NodePosition, imports []*InterpretedString) *ImportsDeclaration {
	return &ImportsDeclaration{
		NodeType:     NodeTypeImportsDeclaration,
		NodePosition: position,
		Imports:      imports,
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
