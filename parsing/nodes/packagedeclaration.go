package nodes

import (
	"fmt"
)

type PackageDeclaration struct {
	NodeType
	NodePosition
	Identifier string
}

func CreatePackageDeclaration(position NodePosition, identifier string) *PackageDeclaration {
	return &PackageDeclaration{
		NodeType:     NodeTypePackageDeclaration,
		NodePosition: position,
		Identifier:   identifier,
	}
}

func (node *PackageDeclaration) String() string {
	return fmt.Sprintf("package %s\n", node.Identifier)
}

func (node *PackageDeclaration) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"package %s\"];\n", node, node.Identifier)
}
