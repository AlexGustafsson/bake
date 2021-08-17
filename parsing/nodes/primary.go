package nodes

type Primary struct {
	NodeType
	NodePosition
	Operand Node
}

type PrimaryOperator int

func CreatePrimary(position NodePosition, operator PrimaryOperator, operand Node) *Primary {
	return &Primary{
		NodeType:     NodeTypePrimary,
		NodePosition: position,
		Operand:      operand,
	}
}

func (node *Primary) String() string {
	return node.Operand.String()
}

func (node *Primary) DotString() string {
	return node.Operand.DotString()
}
