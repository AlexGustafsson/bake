package nodes

type Boolean struct {
	NodeType
	NodePosition
	Value string
}

func CreateBoolean(position NodePosition, value string) *Boolean {
	return &Boolean{
		NodeType:     NodeTypeInteger,
		NodePosition: position,
		Value:        value,
	}
}

func (node *Boolean) String() string {
	return node.Value
}
