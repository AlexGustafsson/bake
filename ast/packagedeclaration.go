package ast

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
