package dot

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func Format(root nodes.Node) string {
	var builder strings.Builder

	switch node := root.(type) {
	case *nodes.AliasDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"alias %s\"];\n", node, node.Identifier)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"expression\"];\n", node, node.Expression)
		builder.WriteString(Format(node.Expression))
	case *nodes.Array:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "array")
		for i, element := range node.Elements {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node, element, i)
			builder.WriteString(Format(element))
		}
	case *nodes.Block:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "block")

		for _, statement := range node.Statements {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, statement)
			builder.WriteString(Format(statement))
		}
	case *nodes.Comparison:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(Format(node.Left))
		builder.WriteString(Format(node.Right))
	case *nodes.Equality:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(Format(node.Left))
		builder.WriteString(Format(node.Right))
	case *nodes.Factor:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(Format(node.Left))
		builder.WriteString(Format(node.Right))
	case *nodes.FunctionDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
		if node.Exported {
			builder.WriteString("export ")
		}
		fmt.Fprintf(&builder, "func %s\"];\n", node.Identifier)

		if node.Signature != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
			builder.WriteString(Format(node.Signature))
		}

		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(Format(node.Block))
	case *nodes.Index:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
		builder.WriteString(Format(node.Operand))
		builder.WriteString(Format(node.Expression))
	case *nodes.Invocation:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "invocation")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		if len(node.Arguments) > 0 {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
			fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
			for i, argument := range node.Arguments {
				fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Arguments, argument, i)
				builder.WriteString(Format(argument))
			}
		}
		builder.WriteString(Format(node.Operand))
	case *nodes.Primary:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "primary")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		builder.WriteString(Format(node.Operand))
	case *nodes.ReturnStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"return\"];\n", node)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"value\"];\n", node, node.Value)
		builder.WriteString(Format(node.Value))
	case *nodes.RuleDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"rule\"];\n", node)

		// Outputs
		fmt.Fprintf(&builder, "\"%p\" [label=\"outputs\"];\n", node.Outputs)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Outputs)
		for i, output := range node.Outputs {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Outputs, output, i)
			builder.WriteString(Format(output))
		}

		// Dependencies
		if len(node.Dependencies) > 0 {
			fmt.Fprintf(&builder, "\"%p\" [label=\"dependencies\"];\n", node.Dependencies)
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Dependencies)
			for i, dependency := range node.Dependencies {
				fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Dependencies, dependency, i)
				builder.WriteString(Format(dependency))
			}
		}

		// Derived
		if node.Derived != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"derived\"];\n", node, node.Derived)
			builder.WriteString(Format(node.Derived))
		}

		// Block
		if node.Block != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
			builder.WriteString(Format(node.Block))
		}
	case *nodes.RuleFunctionDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
		if node.Exported {
			builder.WriteString("export ")
		}
		fmt.Fprintf(&builder, "rule function %s\"];\n", node.Identifier)

		if node.Signature != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
			builder.WriteString(Format(node.Signature))
		}

		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(Format(node.Block))
	case *nodes.Selector:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "selector")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		builder.WriteString(Format(node.Operand))
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	case *nodes.SourceFile:
		builder.WriteString("digraph G {\n")

		fmt.Fprintf(&builder, "\"%p\" [label=\"source file\"];\n", node)

		for _, child := range node.Nodes {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, child)
			builder.WriteString(Format(child))
		}

		builder.WriteString("}\n")
	case *nodes.Term:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(Format(node.Left))
		builder.WriteString(Format(node.Right))
	case *nodes.Unary:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "unary")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Primary, "primary")
		builder.WriteString(Format(node.Primary))
	case *nodes.VariableDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "variable declaration")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
		if node.Expression != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
			builder.WriteString(Format(node.Expression))
		}
	case *nodes.InterpretedString:
		fmt.Fprintf(&builder, "\"%p\" [label=\"interpreted string '%s'\"];\n", node, escape(node.Content))
	case *nodes.RawString:
		fmt.Fprintf(&builder, "\"%p\" [label=\"raw string '%s'\"];\n", node, escape(node.Content))
	case *nodes.Boolean:
		fmt.Fprintf(&builder, "\"%p\" [label=\"boolean '%s'\"];", node, node.Value)
	case *nodes.Comment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"comment '%s'\"];", node, escape(node.Content))
	case *nodes.Identifier:
		fmt.Fprintf(&builder, "\"%p\" [label=\"identifier '%s'\"];\n", node, node.Value)
	case *nodes.ImportSelector:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.From, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.From, node.From)
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	case *nodes.Integer:
		fmt.Fprintf(&builder, "\"%p\" [label=\"integer %s\"];\n", node, node.Value)
	case *nodes.Signature:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "signature")
		fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
		for i, argument := range node.Arguments {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p%d\" [label=\"%d\"];\n", node.Arguments, node.Arguments, i, i)
			fmt.Fprintf(&builder, "\"%p%d\" [label=\"%s\"];\n", node.Arguments, i, argument)
		}
	case *nodes.ShellStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"shell '%s'\"]", node, escape(node.ShellString))
	case *nodes.PackageDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"package declaration '%s'\"];\n", node, node.Identifier)
	case *nodes.ImportsDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "imports")
		for _, literal := range node.Imports {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, literal)
			builder.WriteString(Format(literal))
		}
	case *nodes.Assignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.Increment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "increment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		builder.WriteString(Format(node.Expression))
	case *nodes.Decrement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "decrement")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		builder.WriteString(Format(node.Expression))
	case *nodes.LooseAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "loose assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.AdditionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "addition assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.SubtractionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "subtraction assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.MultiplicationAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "multiplication assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.DivisionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "division assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(Format(node.Expression))
	case *nodes.IfStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"if statement\"];\n", node)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"expression\"];\n", node, node.Expression)
		builder.WriteString(Format(node.Expression))
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"positive branch\"];\n", node, node.PositiveBranch)
		builder.WriteString(Format(node.PositiveBranch))
		if node.NegativeBranch != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"negative branch\"];\n", node, node.NegativeBranch)
			builder.WriteString(Format(node.NegativeBranch))
		}
	default:
		fmt.Fprintf(&builder, "\"%p\" [label=\"UNKNOWN '%s'\"", node, node.Type().String())
	}

	return builder.String()
}

func escape(unescaped string) string {
	escaped := strings.ReplaceAll(unescaped, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "\n", "\\n")
	return escaped
}
