package nodes

import (
	"fmt"
	"strings"
)

type ImportSelector struct {
	NodeType
	Range
	From       string
	Identifier string
}

func CreateImportSelector(r Range, from string, identifier string) *ImportSelector {
	return &ImportSelector{
		NodeType:   NodeTypeImportSelector,
		Range:      r,
		From:       from,
		Identifier: identifier,
	}
}

func (node *ImportSelector) String() string {
	return fmt.Sprintf("%s::%s", node.From, node.Identifier)
}

func (node *ImportSelector) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.From, "operand")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.From, node.From)
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	return builder.String()
}
