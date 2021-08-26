package ast

type FunctionDeclaration struct {
	baseNode
	Exported   bool
	Identifier string
	Signature  *Signature
	Block      *Block
}

func CreateFunctionDeclaration(r *Range, exported bool, identifier string, signature *Signature, block *Block) *FunctionDeclaration {
	return &FunctionDeclaration{
		baseNode: baseNode{
			nodeType:  NodeTypeFunctionDeclaration,
			nodeRange: r,
		},
		Exported:   exported,
		Identifier: identifier,
		// Signature may be nil
		Signature: signature,
		Block:     block,
	}
}
