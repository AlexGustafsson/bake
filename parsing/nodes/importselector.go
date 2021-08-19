package nodes

import (
	"fmt"
	"strings"
)

type ImportSelector struct {
	NodeType
	Range
	Operand    Node
	Identifier string
}

func CreateImportSelector(r Range, operand Node, identifier string) *ImportSelector {
	return &ImportSelector{
		NodeType:   NodeTypeImportSelector,
		Range:      r,
		Operand:    operand,
		Identifier: identifier,
	}
}

func (node *ImportSelector) String() string {
	return fmt.Sprintf("%s::%s", node.Operand.String(), node.Identifier)
}

func (node *ImportSelector) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	builder.WriteString(node.Operand.DotString())
	return builder.String()
}
