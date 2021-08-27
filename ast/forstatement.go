package ast

type ForStatement struct {
	baseNode
	Identifier *Identifier
	Expression Node
	Block      *Block
}

func CreateForStatement(r *Range, identifier *Identifier, expression Node, block *Block) *ForStatement {
	return &ForStatement{
		baseNode: baseNode{
			nodeType:  NodeTypeForStatement,
			nodeRange: r,
		},
		Identifier: identifier,
		Expression: expression,
		Block:      block,
	}
}
