package nodes

import (
	"strings"
)

type RuleFunctionDeclaration struct {
	NodeType
	Range
	Exported   bool
	Identifier string
	// Signature may be nil
	Signature *Signature
	Block     *Block
}

func CreateRuleFunctionDeclaration(r Range, exported bool, identifier string, signature *Signature, block *Block) *RuleFunctionDeclaration {
	return &RuleFunctionDeclaration{
		NodeType:   NodeTypeRuleFunctionDeclaration,
		Range:      r,
		Exported:   exported,
		Identifier: identifier,
		Signature:  signature,
		Block:      block,
	}
}

func (node *RuleFunctionDeclaration) String() string {
	var builder strings.Builder

	if node.Exported {
		builder.WriteString("export ")
	}

	builder.WriteString("rule ")
	builder.WriteString(node.Identifier)
	builder.WriteRune(' ')

	if node.Signature != nil {
		builder.WriteString(node.Signature.String())
		builder.WriteRune(' ')
	}

	builder.WriteString(node.Block.String())

	return builder.String()
}
