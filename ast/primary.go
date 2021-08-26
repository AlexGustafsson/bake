package ast

type Primary struct {
	baseNode
	Operand Node
}

type PrimaryOperator int

func CreatePrimary(r *Range, operator PrimaryOperator, operand Node) *Primary {
	return &Primary{
		baseNode: baseNode{
			nodeType:  NodeTypePrimary,
			nodeRange: r,
		},
		Operand: operand,
	}
}
