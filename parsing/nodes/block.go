package nodes

import (
	"fmt"
	"strings"
)

type Block struct {
	NodeType
	NodePosition
	Statements []Node
}

func CreateBlock(position NodePosition, statements []Node) *Block {
	return &Block{
		NodeType:     NodeTypeBlock,
		NodePosition: position,
		Statements:   statements,
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

	builder.WriteString("}\n")

	return builder.String()
}

func (node *Block) DotString() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "block")

	for _, statement := range node.Statements {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, statement)
		builder.WriteString(statement.DotString())
	}

	return builder.String()
}
