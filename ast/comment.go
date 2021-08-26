package ast

import (
	"fmt"
	"strings"
)

type Comment struct {
	baseNode
	Content string
}

func CreateComment(r *Range, content string) *Comment {
	return &Comment{
		baseNode: baseNode{
			nodeType:  NodeTypeComment,
			nodeRange: r,
		},
		Content: content,
	}
}

func (node *Comment) String() string {
	return fmt.Sprintf("// %s\n", strings.TrimSuffix(node.Content, "\n"))
}
