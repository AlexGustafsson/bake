package nodes

import (
	"strings"
)

type RuleDeclaration struct {
	NodeType
	Range
	Outputs      []Node
	Dependencies []Node
	// May be nil
	Derived Node
	// May be nil if Derived is not nil
	Block *Block
}

func CreateRuleDeclaration(r Range, outputs []Node, dependencies []Node, derived Node, block *Block) *RuleDeclaration {
	return &RuleDeclaration{
		NodeType:     NodeTypeRuleDeclaration,
		Range:        r,
		Outputs:      outputs,
		Dependencies: dependencies,
		Derived:      derived,
		Block:        block,
	}
}

func (node *RuleDeclaration) String() string {
	var builder strings.Builder

	if len(node.Outputs) == 1 {
		builder.WriteString(node.Outputs[0].String())
	} else {
		builder.WriteRune('[')
		for i, output := range node.Outputs {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(output.String())
		}
		builder.WriteRune(']')
	}

	builder.WriteRune(' ')

	if len(node.Dependencies) > 0 {
		builder.WriteRune('[')
		for i, output := range node.Dependencies {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(output.String())
		}
		builder.WriteRune(']')
		builder.WriteRune(' ')
	}

	if node.Derived != nil {
		builder.WriteString(": ")
		builder.WriteString(node.Derived.String())
	}

	if node.Block == nil {
		builder.WriteRune('\n')
	} else {
		builder.WriteRune(' ')
		builder.WriteString(node.Block.String())
		builder.WriteRune('\n')
	}

	return builder.String()
}
