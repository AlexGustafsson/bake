package nodes

import (
	"fmt"
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

	if node.Block != nil {
		builder.WriteRune(' ')
		builder.WriteString(node.Block.String())
	}

	return builder.String()
}

func (node *RuleDeclaration) DotString() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "\"%p\" [label=\"rule\"];\n", node)

	// Outputs
	fmt.Fprintf(&builder, "\"%p\" [label=\"outputs\"];\n", node.Outputs)
	fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Outputs)
	for i, output := range node.Outputs {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Outputs, output, i)
		builder.WriteString(output.DotString())
	}

	// Dependencies
	if len(node.Dependencies) > 0 {
		fmt.Fprintf(&builder, "\"%p\" [label=\"dependencies\"];\n", node.Dependencies)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Dependencies)
		for i, dependency := range node.Dependencies {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Dependencies, dependency, i)
			builder.WriteString(dependency.DotString())
		}
	}

	// Derived
	if node.Derived != nil {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"derived\"];\n", node, node.Derived)
		builder.WriteString(node.Derived.DotString())
	}

	// Block
	if node.Block != nil {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(node.Block.DotString())
	}

	return builder.String()
}
