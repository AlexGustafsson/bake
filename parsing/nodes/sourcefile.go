package nodes

import (
	"fmt"
	"strings"
)

type SourceFile struct {
	NodeType
	NodePosition
	Nodes []Node
}

func CreateSourceFile(position NodePosition) *SourceFile {
	return &SourceFile{
		NodeType:     NodeTypeSourceFile,
		NodePosition: position,
		Nodes:        make([]Node, 0),
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
