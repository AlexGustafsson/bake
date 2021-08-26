package ast

type Index struct {
	baseNode
	Operand    Node
	Expression Node
}

func CreateIndex(r *Range, operand Node, expression Node) *Index {
	return &Index{
		baseNode: baseNode{
			nodeType:  NodeTypeIndex,
			nodeRange: r,
		},
		Operand:    operand,
		Expression: expression,
	}
}
