package ast

type Comparison struct {
	baseNode
	Operator ComparisonOperator
	Left     Node
	Right    Node
}

//go:generate stringer -type=ComparisonOperator
type ComparisonOperator int

const (
	ComparisonOperatorEquals ComparisonOperator = iota
	ComparisonOperatorNotEquals
	ComparisonOperatorLessThan
	ComparisonOperatorLessThanOrEqual
	ComparisonOperatorGreaterThan
	ComparisonOperatorGreaterThanOrEqual
)

func CreateComparison(r *Range, operator ComparisonOperator, left Node, right Node) *Comparison {
	return &Comparison{
		baseNode: baseNode{
			nodeType:  NodeTypeComparison,
			nodeRange: r,
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}
