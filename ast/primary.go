package ast

type Primary struct {
	NodeType
	Range
	Operand Node
}

type PrimaryOperator int

func CreatePrimary(r Range, operator PrimaryOperator, operand Node) *Primary {
	return &Primary{
		NodeType: NodeTypePrimary,
		Range:    r,
		Operand:  operand,
	}
}

func (node *Primary) String() string {
	return node.Operand.String()
}
