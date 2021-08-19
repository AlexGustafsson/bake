package nodes

import (
	"fmt"
	"strings"
)

type SourceFile struct {
	NodeType
	Range
	Nodes []Node
}

func CreateSourceFile(r Range) *SourceFile {
	return &SourceFile{
		NodeType: NodeTypeSourceFile,
		Range:    r,
		Nodes:    make([]Node, 0),
	}
}

func (node *SourceFile) String() string {
	var builder strings.Builder

	for _, node := range node.Nodes {
		builder.WriteString(node.String())
	}

	return builder.String()
}

func (node *SourceFile) DotString() string {
	var builder strings.Builder

	builder.WriteString("digraph G {\n")

	fmt.Fprintf(&builder, "\"%p\" [label=\"source file\"];\n", node)

	for _, child := range node.Nodes {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, child)
		builder.WriteString(child.DotString())
	}

	builder.WriteString("}\n")

	return builder.String()
}
