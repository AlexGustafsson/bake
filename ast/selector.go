package ast

import (
	"fmt"
)

type Selector struct {
	NodeType
	Range
	Operand    Node
	Identifier string
}

func CreateSelector(r Range, operand Node, identifier string) *Selector {
	return &Selector{
		NodeType:   NodeTypeSelector,
		Range:      r,
		Operand:    operand,
		Identifier: identifier,
	}
}

func (node *Selector) String() string {
	return fmt.Sprintf("%s.%s", node.Operand.String(), node.Identifier)
}
