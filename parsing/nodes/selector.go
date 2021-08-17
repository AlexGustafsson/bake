package nodes

import (
	"fmt"
)

type Selector struct {
	NodeType
	NodePosition
	Operand    Node
	Identifier string
}

func CreateSelector(position NodePosition, operand Node, identifier string) *Selector {
	return &Selector{
		NodeType:     NodeTypeSelector,
		NodePosition: position,
		Operand:      operand,
		Identifier:   identifier,
	}
}

func (node *Selector) String() string {
	return fmt.Sprintf("%s.%s", node.Operand.String(), node.Identifier)
}
