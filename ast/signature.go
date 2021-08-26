package ast

type Signature struct {
	baseNode
	Arguments []*Identifier
}

func CreateSignature(r *Range, arguments []*Identifier) *Signature {
	return &Signature{
		baseNode: baseNode{
			nodeType:  NodeTypeSignature,
			nodeRange: r,
		},
		Arguments: arguments,
	}
}
