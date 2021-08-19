package nodes

import (
	"fmt"
	"strings"
)

type Index struct {
	NodeType
	Range
	Operand    Node
	Expression Node
}

func CreateIndex(r Range, operand Node, expression Node) *Index {
	return &Index{
		NodeType:   NodeTypeIndex,
		Range:      r,
		Operand:    operand,
		Expression: expression,
	}
}

func (node *Index) String() string {
	return fmt.Sprintf("%s[%s]", node.Operand.String(), node.Expression)
}

func (node *Index) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
	builder.WriteString(node.Operand.DotString())
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}
