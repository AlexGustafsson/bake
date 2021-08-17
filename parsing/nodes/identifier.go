package nodes

import "fmt"

type Identifier struct {
	NodeType
	NodePosition
	Value string
}

func CreateIdentifier(position NodePosition, value string) *Identifier {
	return &Identifier{
		NodeType:     NodeTypeIdentifier,
		NodePosition: position,
		Value:        value,
	}
}

func (node *Identifier) String() string {
	return node.Value
}

func (node *Identifier) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"%s\"];\n", node, node.Value)
}
