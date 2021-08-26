package ast

type Invocation struct {
	baseNode
	Operand   Node
	Arguments []Node
}

func CreateInvocation(r *Range, operand Node, arguments []Node) *Invocation {
	return &Invocation{
		baseNode: baseNode{
			nodeType:  NodeTypeInvocation,
			nodeRange: r,
		},
		Operand:   operand,
		Arguments: arguments,
	}
}
