package nodes

type Number struct {
	NodeType
	NodePosition
	Value string
}

type Integer struct {
	Number
}

func CreateInteger(position NodePosition, value string) *Integer {
	return &Integer{
		Number{
			NodeType:     NodeTypeInteger,
			NodePosition: position,
			Value:        value,
		},
	}
}

func (node *Number) String() string {
	return node.Value
}
