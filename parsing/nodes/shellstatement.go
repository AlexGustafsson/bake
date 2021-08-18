package nodes

import (
	"fmt"
	"strings"
)

type ShellStatement struct {
	NodeType
	NodePosition
	Multiline   bool
	ShellString string
}

func CreateShellStatement(position NodePosition, multiline bool, shellString string) *ShellStatement {
	return &ShellStatement{
		NodeType:     NodeTypeShellStatement,
		NodePosition: position,
		Multiline:    multiline,
		ShellString:  shellString,
	}
}

func (node *ShellStatement) String() string {
	if node.Multiline {
		return fmt.Sprintf("shell {%s}", node.ShellString)
	} else {
		return fmt.Sprintf("shell %s", node.ShellString)
	}
}

func (node *ShellStatement) DotString() string {
	escaped := strings.ReplaceAll(node.ShellString, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "\n", "\\n")
	return fmt.Sprintf("\"%p\" [label=\"shell '%s'\"]", node, escaped)
}
