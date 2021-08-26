package ast

type RuleFunctionDeclaration struct {
	baseNode
	Exported   bool
	Identifier string
	// Signature may be nil
	Signature *Signature
	Block     *Block
}

func CreateRuleFunctionDeclaration(r *Range, exported bool, identifier string, signature *Signature, block *Block) *RuleFunctionDeclaration {
	return &RuleFunctionDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeRuleFunctionDeclaration,
			nodeRange: r,
		},
		Exported:   exported,
		Identifier: identifier,
		Signature:  signature,
		Block:      block,
	}
}
