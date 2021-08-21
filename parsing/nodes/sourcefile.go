package nodes

import (
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
