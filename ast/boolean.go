package ast

type Boolean struct {
	baseNode
	Value string
}

func CreateBoolean(r *Range, value string) *Boolean {
	return &Boolean{
		baseNode: baseNode{
			nodeType:  NodeTypeInteger,
			nodeRange: r,
		},
		Value: value,
	}
}
