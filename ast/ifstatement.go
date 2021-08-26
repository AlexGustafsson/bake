package ast

type IfStatement struct {
	baseNode
	Expression     Node
	PositiveBranch Node
	// May be nil
	NegativeBranch Node
}

func CreateIfStatement(r *Range, expression Node, positiveBranch Node, negativeBranch Node) *IfStatement {
	return &IfStatement{
		baseNode: baseNode{
			nodeType:  NodeTypeIfStatement,
			nodeRange: r,
		},
		Expression:     expression,
		PositiveBranch: positiveBranch,
		NegativeBranch: negativeBranch,
	}
}
