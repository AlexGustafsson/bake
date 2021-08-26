package ast

type PackageDeclaration struct {
	baseNode
	Identifier string
}

func CreatePackageDeclaration(r *Range, identifier string) *PackageDeclaration {
	return &PackageDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypePackageDeclaration,
			nodeRange: r,
		},
		Identifier: identifier,
	}
}
