package ast

type Term struct {
	baseNode
	Operator AdditiveOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=AdditiveOperator
type AdditiveOperator int

const (
	AdditiveOperatorAddition AdditiveOperator = iota
	AdditiveOperatorSubtraction
)

func CreateTerm(r *Range, operator AdditiveOperator, left Node, right Node) *Term {
	return &Term{
		baseNode: baseNode{
			nodeType:  NodeTypeTerm,
			nodeRange: r,
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}
