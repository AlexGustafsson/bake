package ast

import (
	"fmt"
)

type ReturnStatement struct {
	NodeType
	Range
	Value Node
}

func CreateReturnStatement(r Range, value Node) *ReturnStatement {
	return &ReturnStatement{
		NodeType: NodeTypeReturnStatement,
		Range:    r,
		Value:    value,
	}
}

func (node *ReturnStatement) String() string {
	return fmt.Sprintf("return %s\n", node.Value.String())
}
