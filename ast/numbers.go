package ast

type Integer struct {
	baseNode
	Value string
}

func CreateInteger(r *Range, value string) *Integer {
	return &Integer{
		baseNode: baseNode{
			nodeType:  NodeTypeInteger,
			nodeRange: r,
		},
		Value: value,
	}
}
