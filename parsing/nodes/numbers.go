package nodes

import "fmt"

type Integer struct {
	NodeType
	NodePosition
	Value string
}

func CreateInteger(position NodePosition, value string) *Integer {
	return &Integer{
		NodeType:     NodeTypeInteger,
		NodePosition: position,
		Value:        value,
	}
}

func (node *Integer) String() string {
	return node.Value
}

func (node *Integer) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"integer %s\"];\n", node, node.Value)
}
