package nodes

import "fmt"

type Integer struct {
	NodeType
	Range
	Value string
}

func CreateInteger(r Range, value string) *Integer {
	return &Integer{
		NodeType: NodeTypeInteger,
		Range:    r,
		Value:    value,
	}
}

func (node *Integer) String() string {
	return node.Value
}

func (node *Integer) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"integer %s\"];\n", node, node.Value)
}
