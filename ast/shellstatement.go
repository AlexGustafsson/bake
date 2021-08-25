package ast

import (
	"strings"
)

type ShellStatement struct {
	NodeType
	Range
	Multiline bool
	Parts     []Node
}

func CreateShellStatement(r Range, multiline bool, parts []Node) *ShellStatement {
	return &ShellStatement{
		NodeType:  NodeTypeShellStatement,
		Range:     r,
		Multiline: multiline,
		Parts:     parts,
	}
}

func (node *ShellStatement) String() string {
	var builder strings.Builder
	builder.WriteString("shell ")
	if node.Multiline {
		builder.WriteRune('{')
	}

	for _, part := range node.Parts {
		builder.WriteString(part.String())
	}

	if node.Multiline {
		builder.WriteRune('}')
	}
	builder.WriteRune('\n')
	return builder.String()
}
