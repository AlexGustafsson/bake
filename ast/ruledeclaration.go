package ast

type RuleDeclaration struct {
	baseNode
	Outputs      []Node
	Dependencies []Node
	// May be nil
	Derived Node
	// May be nil if Derived is not nil
	Block *Block
}

func CreateRuleDeclaration(r *Range, outputs []Node, dependencies []Node, derived Node, block *Block) *RuleDeclaration {
	return &RuleDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeRuleDeclaration,
			nodeRange: r,
		},
		Outputs:      outputs,
		Dependencies: dependencies,
		Derived:      derived,
		Block:        block,
	}
}
