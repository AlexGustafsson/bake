package nodes

import (
	"fmt"
)

type ImportSelector struct {
	NodeType
	Range
	From       string
	Identifier string
}

func CreateImportSelector(r Range, from string, identifier string) *ImportSelector {
	return &ImportSelector{
		NodeType:   NodeTypeImportSelector,
		Range:      r,
		From:       from,
		Identifier: identifier,
	}
}

func (node *ImportSelector) String() string {
	return fmt.Sprintf("%s::%s", node.From, node.Identifier)
}
