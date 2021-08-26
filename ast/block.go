package ast

type Block struct {
	baseNode
	Statements []Node
}

func CreateBlock(r *Range, statements []Node) *Block {
	return &Block{
		baseNode: baseNode{
			nodeType:  NodeTypeBlock,
			nodeRange: r,
		},
		Statements: statements,
	}
}
