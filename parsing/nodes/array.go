package nodes

import (
	"fmt"
	"strings"
)

type Array struct {
	NodeType
	Range
	Elements []Node
}

func CreateArray(r Range, elements []Node) *Array {
	return &Array{
		NodeType: NodeTypeComment,
		Range:    r,
		Elements: elements,
	}
}

func (node *Array) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	for i, element := range node.Elements {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(element.String())
	}
	builder.WriteRune(']')
	return builder.String()
}

func (node *Array) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "array")
	for i, element := range node.Elements {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node, element, i)
		builder.WriteString(element.DotString())
	}
	return builder.String()
}
