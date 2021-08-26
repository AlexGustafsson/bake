package ast

type VariableDeclaration struct {
	baseNode
	Identifier string
	Expression Node
}

func CreateVariableDeclaration(r *Range, identifier string, expression Node) *VariableDeclaration {
	return &VariableDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeVariableDeclaration,
			nodeRange: r,
		},
		Identifier: identifier,
		Expression: expression,
	}
}
