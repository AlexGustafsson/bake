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
		NodeType:     NodeTypePackageDeclaration,
		NodePosition: position,
		Content:      content,
	}
}

func (node *Comment) String() string {
	return fmt.Sprintf("// %s\n", strings.TrimSuffix(node.Content, "\n"))
}