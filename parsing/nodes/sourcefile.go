package nodes

import "strings"

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
