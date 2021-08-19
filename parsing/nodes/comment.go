package nodes

import (
	"fmt"
	"strings"
)

type Comment struct {
	NodeType
	Range
	Content string
}

func CreateComment(r Range, content string) *Comment {
	return &Comment{
		NodeType: NodeTypeComment,
		Range:    r,
		Content:  content,
	}
}

func (node *Comment) String() string {
	return fmt.Sprintf("// %s\n", strings.TrimSuffix(node.Content, "\n"))
}

func (node *Comment) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"comment '%s'\"];", node, strings.ReplaceAll(node.Content, "\"", "\\\""))
}
