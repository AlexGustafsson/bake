package ast

import (
	"strings"
)

type Invocation struct {
	NodeType
	Range
	Operand   Node
	Arguments []Node
}

func CreateInvocation(r Range, operand Node, arguments []Node) *Invocation {
	return &Invocation{
		NodeType:  NodeTypeInvocation,
		Range:     r,
		Operand:   operand,
		Arguments: arguments,
	}
}

func (node *Invocation) String() string {
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
