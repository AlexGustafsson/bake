package ast

type Factor struct {
	baseNode
	Operator MultiplicativeOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=MultiplicativeOperator
type MultiplicativeOperator int

const (
	MultiplicativeOperatorMultiplication MultiplicativeOperator = iota
	MultiplicativeOperatorDivision
	MultiplicativeOperatorModulo
)

func CreateFactor(r *Range, operator MultiplicativeOperator, left Node, right Node) *Factor {
	return &Factor{
		baseNode: baseNode{
			nodeType:  NodeTypeFactor,
			nodeRange: r,
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}
