package nodes

type Identifier struct {
	NodeType
	NodePosition
	Value string
}

func CreateIdentifier(position NodePosition, value string) *Identifier {
	return &Identifier{
		NodeType:     NodeTypeIdentifier,
		NodePosition: position,
		Value:        value,
	}
}

func (node *Identifier) String() string {
	return node.Value
}
