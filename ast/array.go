package ast

type Array struct {
	baseNode
	Elements []Node
}

func CreateArray(r *Range, elements []Node) *Array {
	return &Array{
		baseNode: baseNode{
			nodeType:  NodeTypeComment,
			nodeRange: r,
		},
		Elements: elements,
	}
}
