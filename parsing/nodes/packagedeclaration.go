package nodes

import (
	"fmt"
)

type PackageDeclaration struct {
	NodeType
	Range
	Identifier string
}

func CreatePackageDeclaration(r Range, identifier string) *PackageDeclaration {
	return &PackageDeclaration{
		NodeType:   NodeTypePackageDeclaration,
		Range:      r,
		Identifier: identifier,
	}
}

func (node *PackageDeclaration) String() string {
	return fmt.Sprintf("package %s\n", node.Identifier)
}

func (node *PackageDeclaration) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"package declaration '%s'\"];\n", node, node.Identifier)
}
