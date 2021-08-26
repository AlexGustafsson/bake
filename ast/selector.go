package ast

type Selector struct {
	baseNode
	Operand    Node
	Identifier string
}

func CreateSelector(r *Range, operand Node, identifier string) *Selector {
	return &Selector{
		baseNode: baseNode{
			nodeType:  NodeTypeSelector,
			nodeRange: r,
		},
		Operand:    operand,
		Identifier: identifier,
	}
}
