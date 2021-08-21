package nodes

import (
	"fmt"
)

type Index struct {
	NodeType
	Range
	Operand    Node
	Expression Node
}

func CreateIndex(r Range, operand Node, expression Node) *Index {
	return &Index{
		NodeType:   NodeTypeIndex,
		Range:      r,
		Operand:    operand,
		Expression: expression,
	}
}

func (node *Index) String() string {
	return fmt.Sprintf("%s[%s]", node.Operand.String(), node.Expression)
}
