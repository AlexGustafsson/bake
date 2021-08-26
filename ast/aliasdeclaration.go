package ast

type AliasDeclaration struct {
	baseNode
	Exported   bool
	Identifier string
	Expression Node
}

func CreateAliasDeclaration(r *Range, exported bool, identifier string, expression Node) *AliasDeclaration {
	return &AliasDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeAliasDeclaration,
			nodeRange: r,
		},
		Exported:   exported,
		Identifier: identifier,
		Expression: expression,
	}
}
