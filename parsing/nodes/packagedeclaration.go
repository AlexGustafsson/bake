package nodes

import (
	"fmt"

	"github.com/AlexGustafsson/bake/lexing"
)

type PackageDeclaration struct {
	NodeType
	NodePosition
	Identifier lexing.Item
}

func CreatePackageDeclaration(position NodePosition, identifier lexing.Item) *PackageDeclaration {
	return &PackageDeclaration{
		NodeType:     NodeTypePackageDeclaration,
		NodePosition: position,
		Identifier:   identifier,
	}
}

func (node *PackageDeclaration) String() string {
	return fmt.Sprintf("package %s\n", node.Identifier.String())
}
