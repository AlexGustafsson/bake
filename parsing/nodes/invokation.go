package nodes

import "strings"

type Invokation struct {
	NodeType
	NodePosition
	Operand   Node
	Arguments []Node
}

func CreateInvokation(position NodePosition, operand Node, arguments []Node) *Invokation {
	return &Invokation{
		NodeType:     NodeTypeInvokation,
		NodePosition: position,
		Operand:      operand,
		Arguments:    arguments,
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
