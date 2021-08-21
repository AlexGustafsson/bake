package nodes

type Integer struct {
	NodeType
	Range
	Value string
}

func CreateInteger(r Range, value string) *Integer {
	return &Integer{
		NodeType: NodeTypeInteger,
		Range:    r,
		Value:    value,
	}
}

func (node *Integer) String() string {
	return node.Value
}
