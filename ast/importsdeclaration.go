package ast

type ImportsDeclaration struct {
	baseNode
	Imports []*EvaluatedString
}

func CreateImportsDeclaration(r *Range, imports []*EvaluatedString) *ImportsDeclaration {
	return &ImportsDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeImportsDeclaration,
			nodeRange: r,
		},
		Imports: imports,
	}
}
