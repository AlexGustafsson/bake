package nodes

type Boolean struct {
	NodeType
	Range
	Value string
}

func CreateBoolean(r Range, value string) *Boolean {
	return &Boolean{
		NodeType: NodeTypeInteger,
		Range:    r,
		Value:    value,
	}
}

func (node *Boolean) String() string {
	return node.Value
}
