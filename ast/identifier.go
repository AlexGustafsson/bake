package ast

type Identifier struct {
	NodeType
	Range
	Value string
}

func CreateIdentifier(r Range, value string) *Identifier {
	return &Identifier{
		NodeType: NodeTypeIdentifier,
		Range:    r,
		Value:    value,
	}
}

func (node *Identifier) String() string {
	return node.Value
}
