package ast

type Assignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateAssignment(r *Range, expression Node, value Node) *Assignment {
	return &Assignment{
		baseNode: baseNode{
			nodeType:  NodeTypeAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type Increment struct {
	baseNode
	Expression Node
}

func CreateIncrement(r *Range, expression Node) *Increment {
	return &Increment{
		baseNode: baseNode{
			nodeType:  NodeTypeIncrement,
			nodeRange: r,
		},
		Expression: expression,
	}
}

type Decrement struct {
	baseNode
	Expression Node
}

func CreateDecrement(r *Range, expression Node) *Decrement {
	return &Decrement{
		baseNode: baseNode{
			nodeType:  NodeTypeDecrement,
			nodeRange: r,
		},
		Expression: expression,
	}
}

type LooseAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateLooseAssignment(r *Range, expression Node, value Node) *LooseAssignment {
	return &LooseAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeLooseAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type AdditionAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateAdditionAssignment(r *Range, expression Node, value Node) *AdditionAssignment {
	return &AdditionAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeAdditionAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type SubtractionAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateSubtractionAssignment(r *Range, expression Node, value Node) *SubtractionAssignment {
	return &SubtractionAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeSubtractionAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type MultiplicationAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateMultiplicationAssignment(r *Range, expression Node, value Node) *MultiplicationAssignment {
	return &MultiplicationAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeMultiplicationAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type DivisionAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateDivisionAssignment(r *Range, expression Node, value Node) *DivisionAssignment {
	return &DivisionAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeDivisionAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}

type ModuloAssignment struct {
	baseNode
	Expression Node
	Value      Node
}

func CreateModuloAssignment(r *Range, expression Node, value Node) *ModuloAssignment {
	return &ModuloAssignment{
		baseNode: baseNode{
			nodeType:  NodeTypeModuloAssignment,
			nodeRange: r,
		},
		Expression: expression,
		Value:      value,
	}
}
