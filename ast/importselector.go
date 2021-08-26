package ast

type ImportSelector struct {
	baseNode
	From       string
	Identifier string
}

func CreateImportSelector(r *Range, from string, identifier string) *ImportSelector {
	return &ImportSelector{
		baseNode: baseNode{
			nodeType:  NodeTypeImportSelector,
			nodeRange: r,
		},
		From:       from,
		Identifier: identifier,
	}
}
