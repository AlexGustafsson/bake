package ast

type Equality struct {
	baseNode
	Operator EqualityOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=EqualityOperator
type EqualityOperator int

const (
	EqualityOperatorOr EqualityOperator = iota
	EqualityOperatorAnd
)

func CreateEquality(r *Range, operator EqualityOperator, left Node, right Node) *Equality {
	return &Equality{
		baseNode: baseNode{
			nodeType:  NodeTypeEquality,
			nodeRange: r,
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}
