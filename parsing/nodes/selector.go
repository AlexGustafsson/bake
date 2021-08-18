package nodes

import (
	"fmt"
	"strings"
)

type Selector struct {
	NodeType
	NodePosition
	Operand    Node
	Identifier string
}

func CreateSelector(position NodePosition, operand Node, identifier string) *Selector {
	return &Selector{
		NodeType:     NodeTypeSelector,
		NodePosition: position,
		Operand:      operand,
		Identifier:   identifier,
	}
}

func (node *Selector) String() string {
	return fmt.Sprintf("%s.%s", node.Operand.String(), node.Identifier)
}

func (node *Selector) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "selector")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
	builder.WriteString(node.Operand.DotString())
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	return builder.String()
}
