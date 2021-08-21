package nodes

import (
	"fmt"
)

type Assignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateAssignment(r Range, expression Node, value Node) *Assignment {
	return &Assignment{
		NodeType:   NodeTypeAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *Assignment) String() string {
	return fmt.Sprintf("%s = %s", node.Expression.String(), node.Value.String())
}

type Increment struct {
	NodeType
	Range
	Expression Node
}

func CreateIncrement(r Range, expression Node) *Increment {
	return &Increment{
		NodeType:   NodeTypeIncrement,
		Range:      r,
		Expression: expression,
	}
}

func (node *Increment) String() string {
	return fmt.Sprintf("%s++", node.Expression.String())
}

type Decrement struct {
	NodeType
	Range
	Expression Node
}

func CreateDecrement(r Range, expression Node) *Decrement {
	return &Decrement{
		NodeType:   NodeTypeDecrement,
		Range:      r,
		Expression: expression,
	}
}

func (node *Decrement) String() string {
	return fmt.Sprintf("%s--", node.Expression.String())
}

type LooseAssignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateLooseAssignment(r Range, expression Node, value Node) *LooseAssignment {
	return &LooseAssignment{
		NodeType:   NodeTypeLooseAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *LooseAssignment) String() string {
	return fmt.Sprintf("%s ?= %s", node.Expression.String(), node.Value.String())
}

type AdditionAssignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateAdditionAssignment(r Range, expression Node, value Node) *AdditionAssignment {
	return &AdditionAssignment{
		NodeType:   NodeTypeAdditionAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *AdditionAssignment) String() string {
	return fmt.Sprintf("%s += %s", node.Expression.String(), node.Value.String())
}

type SubtractionAssignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateSubtractionAssignment(r Range, expression Node, value Node) *SubtractionAssignment {
	return &SubtractionAssignment{
		NodeType:   NodeTypeSubtractionAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *SubtractionAssignment) String() string {
	return fmt.Sprintf("%s -= %s", node.Expression.String(), node.Value.String())
}

type MultiplicationAssignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateMultiplicationAssignment(r Range, expression Node, value Node) *MultiplicationAssignment {
	return &MultiplicationAssignment{
		NodeType:   NodeTypeMultiplicationAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *MultiplicationAssignment) String() string {
	return fmt.Sprintf("%s *= %s", node.Expression.String(), node.Value.String())
}

type DivisionAssignment struct {
	NodeType
	Range
	Expression Node
	Value      Node
}

func CreateDivisionAssignment(r Range, expression Node, value Node) *DivisionAssignment {
	return &DivisionAssignment{
		NodeType:   NodeTypeDivisionAssignment,
		Range:      r,
		Expression: expression,
		Value:      value,
	}
}

func (node *DivisionAssignment) String() string {
	return fmt.Sprintf("%s /= %s", node.Expression.String(), node.Value.String())
}
