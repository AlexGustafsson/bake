package ast

type Unary struct {
	baseNode
	Operator UnaryOperator
	Primary  Node
}

//go:generate stringer -type=UnaryOperator
type UnaryOperator int

const (
	UnaryOperatorSubtraction UnaryOperator = iota
	UnaryOperatorNot
	UnaryOperatorSpread
)

func CreateUnary(r *Range, operator UnaryOperator, primary Node) *Unary {
	return &Unary{
		baseNode: baseNode{
			nodeType:  NodeTypeUnary,
			nodeRange: r,
		},
		Operator: operator,
		Primary:  primary,
	}
}
