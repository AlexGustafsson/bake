package nodes

import (
	"fmt"
)

type Index struct {
	NodeType
	NodePosition
	Operand    Node
	Expression Node
}

func CreateIndex(position NodePosition, operand Node, expression Node) *Index {
	return &Index{
		NodeType:     NodeTypeIndex,
		NodePosition: position,
		Operand:      operand,
		Expression:   expression,
	}
}

func (node *Index) String() string {
	return fmt.Sprintf("%s[%s]", node.Operand.String(), node.Expression)
}
