package nodes

import "fmt"

type Identifier struct {
	NodeType
	Range
	Value string
}

func CreateIdentifier(r Range, value string) *Identifier {
	return &Identifier{
		NodeType: NodeTypeIdentifier,
		Range:    r,
		Value:    value,
	}
}

func (node *Identifier) String() string {
	return node.Value
}

func (node *Identifier) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"identifier '%s'\"];\n", node, node.Value)
}
