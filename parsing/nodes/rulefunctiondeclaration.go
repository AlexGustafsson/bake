package nodes

import (
	"fmt"
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

func (node *RuleFunctionDeclaration) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
	if node.Exported {
		builder.WriteString("export ")
	}
	fmt.Fprintf(&builder, "rule function %s\"];\n", node.Identifier)

	if node.Signature != nil {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
		builder.WriteString(node.Signature.DotString())
	}

	fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
	builder.WriteString(node.Block.DotString())
	return builder.String()
}
