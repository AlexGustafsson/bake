package nodes

import (
	"fmt"
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
	return fmt.Sprintf("// %s\n", node.Content)
}
