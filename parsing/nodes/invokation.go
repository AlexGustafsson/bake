package nodes

import (
	"fmt"
	"strings"
)

type Invokation struct {
	NodeType
	Range
	Operand   Node
	Arguments []Node
}

func CreateInvokation(r Range, operand Node, arguments []Node) *Invokation {
	return &Invokation{
		NodeType:  NodeTypeInvokation,
		Range:     r,
		Operand:   operand,
		Arguments: arguments,
	}
}

func (node *Invokation) String() string {
	var builder strings.Builder

	builder.WriteString(node.Operand.String())
	builder.WriteRune('(')

	for i, argument := range node.Arguments {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(argument.String())
	}

	builder.WriteRune(')')

	return builder.String()
}

func (node *Invokation) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "invocation")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
	if len(node.Arguments) > 0 {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
		fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
		for i, argument := range node.Arguments {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p%d\"\n", node.Arguments, &argument, i)
			fmt.Fprintf(&builder, "\"%p%d\" [label=\"%s\"];\n", &argument, i, argument)
		}
	}
	builder.WriteString(node.Operand.DotString())
	return builder.String()
}
