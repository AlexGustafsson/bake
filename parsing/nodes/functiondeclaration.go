package nodes

import (
	"fmt"
	"strings"
)

type FunctionDeclaration struct {
	NodeType
	Range
	Exported   bool
	Identifier string
	Signature  *Signature
	Block      *Block
}

func CreateFunctionDeclaration(r Range, exported bool, identifier string, signature *Signature, block *Block) *FunctionDeclaration {
	return &FunctionDeclaration{
		NodeType:   NodeTypeFunctionDeclaration,
		Range:      r,
		Exported:   exported,
		Identifier: identifier,
		// Signature may be nil
		Signature: signature,
		Block:     block,
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
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
	if node.Exported {
		builder.WriteString("export ")
	}
	fmt.Fprintf(&builder, "func %s\"];\n", node.Identifier)

	if node.Signature != nil {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
		builder.WriteString(node.Signature.DotString())
	}

	fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
	builder.WriteString(node.Block.DotString())
	return builder.String()
}
