package ast

import (
	"fmt"
)

type ShellStatement struct {
	NodeType
	Range
	Multiline   bool
	ShellString string
}

func CreateShellStatement(r Range, multiline bool, shellString string) *ShellStatement {
	return &ShellStatement{
		NodeType:    NodeTypeShellStatement,
		Range:       r,
		Multiline:   multiline,
		ShellString: shellString,
	}
}

func (node *ShellStatement) String() string {
	if node.Multiline {
		return fmt.Sprintf("shell {%s}", node.ShellString)
	} else {
		return fmt.Sprintf("shell %s", node.ShellString)
	}
}