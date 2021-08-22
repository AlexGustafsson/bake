package ast

import (
	"strings"
)

type Block struct {
	NodeType
	Range
	Statements []Node
}

func CreateBlock(r Range, statements []Node) *Block {
	return &Block{
		NodeType:   NodeTypeBlock,
		Range:      r,
		Statements: statements,
	}
}

func (node *Block) String() string {
	var builder strings.Builder

	builder.WriteString("{\n")

	// TODO: Fix indent
	for _, statement := range node.Statements {
		builder.WriteString("  ")
		builder.WriteString(statement.String())
		builder.WriteRune('\n')
	}

	builder.WriteString("}")

	return builder.String()
}
