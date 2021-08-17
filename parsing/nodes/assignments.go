package nodes

import (
	"fmt"
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
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
	return fmt.Sprintf("\"%p\" [label=\"increment\"];\n\"%p\" -> \"%p\";\n%s", node, node, node.Expression, node.Expression.DotString())
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
	return fmt.Sprintf("\"%p\" [label=\"decrement\"];\n\"%p\" -> \"%p\";\n%s", node, node, node.Expression, node.Expression.DotString())
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
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
	return fmt.Sprintf("\"%p\" -> \"%p\" [label=\"%s\"];\n\"%p\" -> \"%p\" [label=\"%s\"];\n%s%s", node, node.Expression, "left", node, node.Value, "right", node.Expression.DotString(), node.Value.DotString())
}
