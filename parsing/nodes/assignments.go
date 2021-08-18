package nodes

import (
	"fmt"
	"strings"
)

type Assignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateAssignment(position NodePosition, expression Node, value Node) *Assignment {
	return &Assignment{
		NodeType:     NodeTypeAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *Assignment) String() string {
	return fmt.Sprintf("%s = %s", node.Expression.String(), node.Value.String())
}

func (node *Assignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type Increment struct {
	NodeType
	NodePosition
	Expression Node
}

func CreateIncrement(position NodePosition, expression Node) *Increment {
	return &Increment{
		NodeType:     NodeTypeIncrement,
		NodePosition: position,
		Expression:   expression,
	}
}

func (node *Increment) String() string {
	return fmt.Sprintf("%s++", node.Expression.String())
}

func (node *Increment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "increment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type Decrement struct {
	NodeType
	NodePosition
	Expression Node
}

func CreateDecrement(position NodePosition, expression Node) *Decrement {
	return &Decrement{
		NodeType:     NodeTypeDecrement,
		NodePosition: position,
		Expression:   expression,
	}
}

func (node *Decrement) String() string {
	return fmt.Sprintf("%s--", node.Expression.String())
}

func (node *Decrement) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "decrement")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type LooseAssignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateLooseAssignment(position NodePosition, expression Node, value Node) *LooseAssignment {
	return &LooseAssignment{
		NodeType:     NodeTypeLooseAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *LooseAssignment) String() string {
	return fmt.Sprintf("%s ?= %s", node.Expression.String(), node.Value.String())
}

func (node *LooseAssignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "loose assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type AdditionAssignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateAdditionAssignment(position NodePosition, expression Node, value Node) *AdditionAssignment {
	return &AdditionAssignment{
		NodeType:     NodeTypeAdditionAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *AdditionAssignment) String() string {
	return fmt.Sprintf("%s += %s", node.Expression.String(), node.Value.String())
}

func (node *AdditionAssignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "addition assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type SubtractionAssignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateSubtractionAssignment(position NodePosition, expression Node, value Node) *SubtractionAssignment {
	return &SubtractionAssignment{
		NodeType:     NodeTypeSubtractionAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *SubtractionAssignment) String() string {
	return fmt.Sprintf("%s -= %s", node.Expression.String(), node.Value.String())
}

func (node *SubtractionAssignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "subtraction assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type MultiplicationAssignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateMultiplicationAssignment(position NodePosition, expression Node, value Node) *MultiplicationAssignment {
	return &MultiplicationAssignment{
		NodeType:     NodeTypeMultiplicationAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *MultiplicationAssignment) String() string {
	return fmt.Sprintf("%s *= %s", node.Expression.String(), node.Value.String())
}

func (node *MultiplicationAssignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "multiplication assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}

type DivisionAssignment struct {
	NodeType
	NodePosition
	Expression Node
	Value      Node
}

func CreateDivisionAssignment(position NodePosition, expression Node, value Node) *DivisionAssignment {
	return &DivisionAssignment{
		NodeType:     NodeTypeDivisionAssignment,
		NodePosition: position,
		Expression:   expression,
		Value:        value,
	}
}

func (node *DivisionAssignment) String() string {
	return fmt.Sprintf("%s /= %s", node.Expression.String(), node.Value.String())
}

func (node *DivisionAssignment) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "division assignment")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
	builder.WriteString(node.Expression.DotString())
	return builder.String()
}
