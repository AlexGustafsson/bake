package nodes

import "fmt"

type Boolean struct {
	NodeType
	NodePosition
	Value string
}

func CreateBoolean(position NodePosition, value string) *Boolean {
	return &Boolean{
		NodeType:     NodeTypeInteger,
		NodePosition: position,
		Value:        value,
	}
}

func (node *Boolean) String() string {
	return node.Value
}

func (node *Boolean) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"%s\"];", node, node.Value)
}
