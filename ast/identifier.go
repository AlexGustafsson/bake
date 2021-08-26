package ast

type Identifier struct {
	baseNode
	Value string
}

func CreateIdentifier(r *Range, value string) *Identifier {
	return &Identifier{
		baseNode: baseNode{
			nodeType:  NodeTypeIdentifier,
			nodeRange: r,
		},
		Value: value,
	}
}
