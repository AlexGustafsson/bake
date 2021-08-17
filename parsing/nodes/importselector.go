package nodes

import (
	"fmt"
)

type ImportSelector struct {
	NodeType
	NodePosition
	Operand    Node
	Identifier string
}

func CreateImportSelector(position NodePosition, operand Node, identifier string) *ImportSelector {
	return &ImportSelector{
		NodeType:     NodeTypeImportSelector,
		NodePosition: position,
		Operand:      operand,
		Identifier:   identifier,
	}
}

func (node *ImportSelector) String() string {
	return fmt.Sprintf("%s::%s", node.Operand.String(), node.Identifier)
}
