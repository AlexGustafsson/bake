package dot

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/ast"
)

func FormatTree(root ast.Node) string {
	var builder strings.Builder

	builder.WriteString("digraph G {\n")

	builder.WriteString("\"start\" [label=\"source\"];\n")
	fmt.Fprintf(&builder, "\"start\" -> \"%p\";\n", root)

	builder.WriteString(formatTree(root))

	builder.WriteString("}\n")
	return builder.String()
}

func formatTree(root ast.Node) string {
	var builder strings.Builder
	switch node := root.(type) {
	case *ast.AliasDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
		if node.Exported {
			builder.WriteString("export ")
		}
		fmt.Fprintf(&builder, "alias %s\"];\n", node.Identifier)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"expression\"];\n", node, node.Expression)
		builder.WriteString(formatTree(node.Expression))
	case *ast.Array:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "array")
		for i, element := range node.Elements {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node, element, i)
			builder.WriteString(formatTree(element))
		}
	case *ast.Block:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "block")

		for _, statement := range node.Statements {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, statement)
			builder.WriteString(formatTree(statement))
		}
	case *ast.Comparison:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(formatTree(node.Left))
		builder.WriteString(formatTree(node.Right))
	case *ast.Equality:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(formatTree(node.Left))
		builder.WriteString(formatTree(node.Right))
	case *ast.Factor:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(formatTree(node.Left))
		builder.WriteString(formatTree(node.Right))
	case *ast.FunctionDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
		if node.Exported {
			builder.WriteString("export ")
		}
		fmt.Fprintf(&builder, "func %s\"];\n", node.Identifier)

		if node.Signature != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
			builder.WriteString(formatTree(node.Signature))
		}

		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(formatTree(node.Block))
	case *ast.Index:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
		builder.WriteString(formatTree(node.Operand))
		builder.WriteString(formatTree(node.Expression))
	case *ast.Invocation:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "invocation")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		if len(node.Arguments) > 0 {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
			fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
			for i, argument := range node.Arguments {
				fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Arguments, argument, i)
				builder.WriteString(formatTree(argument))
			}
		}
		builder.WriteString(formatTree(node.Operand))
	case *ast.Primary:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "primary")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		builder.WriteString(formatTree(node.Operand))
	case *ast.ReturnStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"return\"];\n", node)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"value\"];\n", node, node.Value)
		builder.WriteString(formatTree(node.Value))
	case *ast.RuleDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"rule\"];\n", node)

		// Outputs
		fmt.Fprintf(&builder, "\"%p\" [label=\"outputs\"];\n", node.Outputs)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Outputs)
		for i, output := range node.Outputs {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Outputs, output, i)
			builder.WriteString(formatTree(output))
		}

		// Dependencies
		if len(node.Dependencies) > 0 {
			fmt.Fprintf(&builder, "\"%p\" [label=\"dependencies\"];\n", node.Dependencies)
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Dependencies)
			for i, dependency := range node.Dependencies {
				fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Dependencies, dependency, i)
				builder.WriteString(formatTree(dependency))
			}
		}

		// Derived
		if node.Derived != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"derived\"];\n", node, node.Derived)
			builder.WriteString(formatTree(node.Derived))
		}

		// Block
		if node.Block != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
			builder.WriteString(formatTree(node.Block))
		}
	case *ast.RuleFunctionDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"", node)
		if node.Exported {
			builder.WriteString("export ")
		}
		fmt.Fprintf(&builder, "rule function %s\"];\n", node.Identifier)

		if node.Signature != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"signature\"];\n", node, node.Signature)
			builder.WriteString(formatTree(node.Signature))
		}

		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(formatTree(node.Block))
	case *ast.Selector:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "selector")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Operand, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		builder.WriteString(formatTree(node.Operand))
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	case *ast.Term:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, node.Operator.String())
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Left, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Right, "right")
		builder.WriteString(formatTree(node.Left))
		builder.WriteString(formatTree(node.Right))
	case *ast.Unary:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "unary")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Primary, "primary")
		builder.WriteString(formatTree(node.Primary))
	case *ast.VariableDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "variable declaration")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
		if node.Expression != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "expression")
			builder.WriteString(formatTree(node.Expression))
		}
	case *ast.EvaluatedString:
		fmt.Fprintf(&builder, "\"%p\" [label=\"evaluated string\"];\n", node)
		for i, part := range node.Parts {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node, part, i)
			builder.WriteString(formatTree(part))
		}
	case *ast.StringPart:
		fmt.Fprintf(&builder, "\"%p\" [label=\"string part '%s'\"];\n", node, escape(node.Content))
	case *ast.RawString:
		fmt.Fprintf(&builder, "\"%p\" [label=\"raw string '%s'\"];\n", node, escape(node.Content))
	case *ast.Boolean:
		fmt.Fprintf(&builder, "\"%p\" [label=\"boolean '%s'\"];", node, node.Value)
	case *ast.Comment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"comment '%s'\"];", node, escape(node.Content))
	case *ast.Identifier:
		fmt.Fprintf(&builder, "\"%p\" [label=\"identifier '%s'\"];\n", node, node.Value)
	case *ast.ImportSelector:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "index")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.From, "operand")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, &node.Identifier, "identifier")
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.From, node.From)
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", &node.Identifier, node.Identifier)
	case *ast.Integer:
		fmt.Fprintf(&builder, "\"%p\" [label=\"integer %s\"];\n", node, node.Value)
	case *ast.Signature:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "signature")
		fmt.Fprintf(&builder, "\"%p\" [label=\"arguments\"];\n", node.Arguments)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Arguments)
		for i, argument := range node.Arguments {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node.Arguments, argument, i)
			fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", argument, argument.Value)
		}
	case *ast.ShellStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"shell\"];\n", node)
		for i, part := range node.Parts {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%d\"];\n", node, part, i)
			builder.WriteString(formatTree(part))
		}
	case *ast.PackageDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"package declaration '%s'\"];\n", node, node.Identifier)
	case *ast.ImportsDeclaration:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "imports")
		for _, literal := range node.Imports {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, literal)
			builder.WriteString(formatTree(literal))
		}
	case *ast.Assignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.Increment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "increment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		builder.WriteString(formatTree(node.Expression))
	case *ast.Decrement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "decrement")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		builder.WriteString(formatTree(node.Expression))
	case *ast.LooseAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "loose assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.AdditionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "addition assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.SubtractionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "subtraction assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.MultiplicationAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "multiplication assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.DivisionAssignment:
		fmt.Fprintf(&builder, "\"%p\" [label=\"%s\"];\n", node, "division assignment")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Expression, "left")
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"%s\"];\n", node, node.Value, "right")
		builder.WriteString(formatTree(node.Expression))
		builder.WriteString(formatTree(node.Value))
	case *ast.IfStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"if statement\"];\n", node)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"expression\"];\n", node, node.Expression)
		builder.WriteString(formatTree(node.Expression))
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"positive branch\"];\n", node, node.PositiveBranch)
		builder.WriteString(formatTree(node.PositiveBranch))
		if node.NegativeBranch != nil {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"negative branch\"];\n", node, node.NegativeBranch)
			builder.WriteString(formatTree(node.NegativeBranch))
		}
	case *ast.ForStatement:
		fmt.Fprintf(&builder, "\"%p\" [label=\"for statement\"];\n", node)
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"identifier\"];\n", node, node.Identifier)
		builder.WriteString(formatTree(node.Identifier))
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\" [label=\"expression\"];\n", node, node.Expression)
		builder.WriteString(formatTree(node.Expression))
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, node.Block)
		builder.WriteString(formatTree(node.Block))
	case *ast.Object:
		fmt.Fprintf(&builder, "\"%p\" [label=\"object\"];\n", node)
		for key, value := range node.Pairs {
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", node, key)
			builder.WriteString(formatTree(key))
			fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", key, value)
			builder.WriteString(formatTree(value))
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
