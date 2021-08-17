package nodes

import (
	"strings"
)

type FunctionDeclaration struct {
	NodeType
	NodePosition
	Exported   bool
	Identifier string
	Signature  *Signature
	Block      *Block
}

func CreateFunctionDeclaration(position NodePosition, exported bool, identifier string, signature *Signature, block *Block) *FunctionDeclaration {
	return &FunctionDeclaration{
		NodeType:     NodeTypeFunctionDeclaration,
		NodePosition: position,
		Exported:     exported,
		Identifier:   identifier,
		Signature:    signature,
		Block:        block,
	}
}

func (node *FunctionDeclaration) String() string {
	var builder strings.Builder

	if node.Exported {
		builder.WriteString("export ")
	}

	builder.WriteString("func ")
	builder.WriteString(node.Identifier)
	builder.WriteRune(' ')

	if node.Signature != nil {
		builder.WriteString(node.Signature.String())
		builder.WriteRune(' ')
	}

	builder.WriteString(node.Block.String())

	return builder.String()
}

func (node *FunctionDeclaration) DotString() string {
	return ""
}
