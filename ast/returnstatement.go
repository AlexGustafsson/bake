package ast

type ReturnStatement struct {
	baseNode
	Value Node
}

func CreateReturnStatement(r *Range, value Node) *ReturnStatement {
	return &ReturnStatement{
		baseNode: baseNode{
			nodeType:  NodeTypeReturnStatement,
			nodeRange: r,
		},
		Value: value,
	}
}
