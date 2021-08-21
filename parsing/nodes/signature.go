package nodes

import (
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
