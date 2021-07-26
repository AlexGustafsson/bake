package nodes

import "strings"

type SourceFile struct {
	NodeType
	NodePosition
	PackageDeclaration   *PackageDeclaration
	TopLevelDeclarations []Node
}

func CreateSourceFile(position NodePosition) *SourceFile {
	return &SourceFile{
		NodeType:             NodeTypeSourceFile,
		NodePosition:         position,
		PackageDeclaration:   nil,
		TopLevelDeclarations: make([]Node, 0),
	}
}

func (node *SourceFile) String() string {
	var builder strings.Builder

	if node.PackageDeclaration != nil {
		builder.WriteString(node.PackageDeclaration.String())
	}

	for _, node := range node.TopLevelDeclarations {
		builder.WriteString(node.String())
	}

	return builder.String()
}
