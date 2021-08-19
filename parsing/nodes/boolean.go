package nodes

import "fmt"

type Boolean struct {
	NodeType
	Range
	Value string
}

func CreateBoolean(r Range, value string) *Boolean {
	return &Boolean{
		NodeType: NodeTypeInteger,
		Range:    r,
		Value:    value,
	}
}

func (node *Boolean) String() string {
	return node.Value
}

func (node *Boolean) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"boolean '%s'\"];", node, node.Value)
}
