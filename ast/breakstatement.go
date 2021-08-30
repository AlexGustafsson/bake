package ast

type BreakStatement struct {
	baseNode
	Value Node
}

func CreateBreakStatement(r *Range) *BreakStatement {
	return &BreakStatement{
		baseNode: baseNode{
			nodeType:  NodeTypeBreakStatement,
			nodeRange: r,
		},
	}
}
