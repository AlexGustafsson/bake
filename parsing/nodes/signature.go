package nodes

import (
	"fmt"
	"strings"
)

type Signature struct {
	NodeType
	Range
	Arguments []string
}

func CreateSignature(r Range, arguments []string) *Signature {
	return &Signature{
		NodeType:  NodeTypeSignature,
		Range:     r,
		Arguments: arguments,
	}
}

func (node *Signature) String() string {
	var builder strings.Builder

	if len(node.Arguments) > 0 {

		builder.WriteRune('(')

		for i, argument := range node.Arguments {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(argument)
		}

		builder.WriteRune(')')

	}

	return builder.String()
}

func (node *Signature) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "signature")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
	fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
	for i, argument := range node.Arguments {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p%d\"\n", node.Arguments, &argument, i)
		fmt.Fprintf(&builder, "\"%p%d\" [label=\"%s\"];\n", &argument, i, argument)
	}
	return builder.String()
}
