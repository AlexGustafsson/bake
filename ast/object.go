package ast

type Object struct {
	baseNode
	Pairs map[*Identifier]Node
}

func CreateObject(r *Range, pairs map[*Identifier]Node) *Object {
	return &Object{
		baseNode: baseNode{
			nodeType:  NodeTypeObject,
			nodeRange: r,
		},
		Pairs: pairs,
	}
}
