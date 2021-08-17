package nodes

import (
	"fmt"
	"strings"
)

type Comment struct {
	NodeType
	NodePosition
	Content string
}

func CreateComment(position NodePosition, content string) *Comment {
	return &Comment{
		NodeType:     NodeTypeComment,
		NodePosition: position,
		Content:      content,
	}
}

func (node *Comment) String() string {
	return fmt.Sprintf("// %s\n", strings.TrimSuffix(node.Content, "\n"))
}

func (node *Comment) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"%s\"];", node, node.Content)
}
