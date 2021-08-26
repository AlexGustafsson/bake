package formatting

// func (node *ast.AliasDeclaration) String() string {
// 	var builder strings.Builder

// 	if node.Exported {
// 		builder.WriteString("export ")
// 	}

// 	fmt.Fprintf(&builder, "alias %s : ", node.Identifier)
// 	builder.WriteString(node.Expression.String())
// 	builder.WriteRune('\n')

// 	return builder.String()
// }

// func (node *ast.Array) String() string {
// 	var builder strings.Builder
// 	builder.WriteRune('[')
// 	for i, element := range node.Elements {
// 		if i > 0 {
// 			builder.WriteString(", ")
// 		}
// 		builder.WriteString(element.String())
// 	}
// 	builder.WriteRune(']')
// 	return builder.String()
// }

// func (node *ast.Assignment) String() string {
// 	return fmt.Sprintf("%s = %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.Increment) String() string {
// 	return fmt.Sprintf("%s++", node.Expression.String())
// }

// func (node *ast.Decrement) String() string {
// 	return fmt.Sprintf("%s--", node.Expression.String())
// }

// func (node *ast.LooseAssignment) String() string {
// 	return fmt.Sprintf("%s ?= %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.AdditionAssignment) String() string {
// 	return fmt.Sprintf("%s += %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.SubtractionAssignment) String() string {
// 	return fmt.Sprintf("%s -= %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.MultiplicationAssignment) String() string {
// 	return fmt.Sprintf("%s *= %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.DivisionAssignment) String() string {
// 	return fmt.Sprintf("%s /= %s", node.Expression.String(), node.Value.String())
// }

// func (node *ast.Block) String() string {
// 	var builder strings.Builder

// 	builder.WriteString("{\n")

// 	// TODO: Fix indent
// 	for _, statement := range node.Statements {
// 		builder.WriteString("  ")
// 		builder.WriteString(statement.String())
// 		builder.WriteRune('\n')
// 	}

// 	builder.WriteString("}")

// 	return builder.String()
// }

// func (node *ast.Boolean) String() string {
// 	return node.Value
// }

// func (node *ast.Comparison) String() string {
// 	var builder strings.Builder

// 	builder.WriteString(node.Left.String())

// 	builder.WriteByte(' ')

// 	switch node.Operator {
// 	case ComparisonOperatorEquals:
// 		builder.WriteString("==")
// 	case ComparisonOperatorNotEquals:
// 		builder.WriteString("!=")
// 	case ComparisonOperatorLessThan:
// 		builder.WriteRune('<')
// 	case ComparisonOperatorLessThanOrEqual:
// 		builder.WriteString("<=")
// 	case ComparisonOperatorGreaterThan:
// 		builder.WriteRune('>')
// 	case ComparisonOperatorGreaterThanOrEqual:
// 		builder.WriteString(">=")
// 	}

// 	builder.WriteByte(' ')

// 	builder.WriteString(node.Right.String())

// 	return builder.String()
// }

// func (node *ast.Equality) String() string {
// 	var builder strings.Builder

// 	builder.WriteString(node.Left.String())

// 	builder.WriteByte(' ')

// 	switch node.Operator {
// 	case EqualityOperatorOr:
// 		builder.WriteString("||")
// 	case EqualityOperatorAnd:
// 		builder.WriteString("&&")
// 	}

// 	builder.WriteByte(' ')

// 	builder.WriteString(node.Right.String())

// 	return builder.String()
// }

// func (node *ast.EvaluatedString) String() string {
// 	var builder strings.Builder

// 	builder.WriteRune('"')

// 	for _, part := range node.Parts {
// 		builder.WriteString(part.String())
// 	}

// 	builder.WriteRune('"')

// 	return builder.String()
// }

// func (node *ast.StringPart) String() string {
// 	return node.Content
// }

// func (node *ast.Factor) String() string {
// 	var builder strings.Builder

// 	builder.WriteString(node.Left.String())

// 	builder.WriteByte(' ')

// 	switch node.Operator {
// 	case MultiplicativeOperatorMultiplication:
// 		builder.WriteString("*")
// 	case MultiplicativeOperatorDivision:
// 		builder.WriteString("/")
// 	}

// 	builder.WriteByte(' ')

// 	builder.WriteString(node.Right.String())

// 	return builder.String()
// }

// func (node *ast.FunctionDeclaration) String() string {
// 	var builder strings.Builder

// 	if node.Exported {
// 		builder.WriteString("export ")
// 	}

// 	builder.WriteString("func ")
// 	builder.WriteString(node.Identifier)
// 	builder.WriteRune(' ')

// 	if node.Signature != nil {
// 		builder.WriteString(node.Signature.String())
// 		builder.WriteRune(' ')
// 	}

// 	builder.WriteString(node.Block.String())
// 	builder.WriteRune('\n')

// 	return builder.String()
// }

// func (node *ast.Identifier) String() string {
// 	return node.Value
// }

// func (node *ast.IfStatement) String() string {
// 	var builder strings.Builder
// 	fmt.Fprintf(&builder, "if %s ", node.Expression.String())
// 	builder.WriteString(node.PositiveBranch.String())
// 	if node.NegativeBranch != nil {
// 		fmt.Fprintf(&builder, " else %s", node.NegativeBranch.String())
// 	}
// 	return builder.String()
// }

// func (node *ast.ImportsDeclaration) String() string {
// 	var builder strings.Builder

// 	builder.WriteString("import (\n")

// 	for _, node := range node.Imports {
// 		builder.WriteString(node.String())
// 		builder.WriteRune('\n')
// 	}

// 	builder.WriteString(")\n")

// 	return builder.String()
// }

// func (node *ast.ImportSelector) String() string {
// 	return fmt.Sprintf("%s::%s", node.From, node.Identifier)
// }

// func (node *ast.Index) String() string {
// 	return fmt.Sprintf("%s[%s]", node.Operand.String(), node.Expression)
// }

// func (node *ast.Invocation) String() string {
// 	var builder strings.Builder

// 	builder.WriteString(node.Operand.String())
// 	builder.WriteRune('(')

// 	for i, argument := range node.Arguments {
// 		if i > 0 {
// 			builder.WriteString(", ")
// 		}
// 		builder.WriteString(argument.String())
// 	}

// 	builder.WriteRune(')')

// 	return builder.String()
// }

// func (node *ast.Integer) String() string {
// 	return node.Value
// }

// func (node *ast.PackageDeclaration) String() string {
// 	return fmt.Sprintf("package %s\n", node.Identifier)
// }

// func (node *ast.Primary) String() string {
// 	return node.Operand.String()
// }

// func (node *ast.ReturnStatement) String() string {
// 	return fmt.Sprintf("return %s\n", node.Value.String())
// }

// func (node *ast.RuleDeclaration) String() string {
// 	var builder strings.Builder

// 	if len(node.Outputs) == 1 {
// 		builder.WriteString(node.Outputs[0].String())
// 	} else {
// 		builder.WriteRune('[')
// 		for i, output := range node.Outputs {
// 			if i > 0 {
// 				builder.WriteString(", ")
// 			}
// 			builder.WriteString(output.String())
// 		}
// 		builder.WriteRune(']')
// 	}

// 	builder.WriteRune(' ')

// 	if len(node.Dependencies) > 0 {
// 		builder.WriteRune('[')
// 		for i, output := range node.Dependencies {
// 			if i > 0 {
// 				builder.WriteString(", ")
// 			}
// 			builder.WriteString(output.String())
// 		}
// 		builder.WriteRune(']')
// 		builder.WriteRune(' ')
// 	}

// 	if node.Derived != nil {
// 		builder.WriteString(": ")
// 		builder.WriteString(node.Derived.String())
// 	}

// 	if node.Block == nil {
// 		builder.WriteRune('\n')
// 	} else {
// 		builder.WriteRune(' ')
// 		builder.WriteString(node.Block.String())
// 		builder.WriteRune('\n')
// 	}

// 	return builder.String()
// }

// func (node *ast.RuleFunctionDeclaration) String() string {
// 	var builder strings.Builder

// 	if node.Exported {
// 		builder.WriteString("export ")
// 	}

// 	builder.WriteString("rule ")
// 	builder.WriteString(node.Identifier)
// 	builder.WriteRune(' ')

// 	if node.Signature != nil {
// 		builder.WriteString(node.Signature.String())
// 		builder.WriteRune(' ')
// 	}

// 	builder.WriteString(node.Block.String())
// 	builder.WriteRune('\n')

// 	return builder.String()
// }

// func (node *ast.Selector) String() string {
// 	return fmt.Sprintf("%s.%s", node.Operand.String(), node.Identifier)
// }

// func (node *ast.ShellStatement) String() string {
// 	var builder strings.Builder
// 	builder.WriteString("shell ")
// 	if node.Multiline {
// 		builder.WriteRune('{')
// 	}

// 	for _, part := range node.Parts {
// 		builder.WriteString(part.String())
// 	}

// 	if node.Multiline {
// 		builder.WriteRune('}')
// 	}
// 	builder.WriteRune('\n')
// 	return builder.String()
// }

// func (node *ast.Signature) String() string {
// 	var builder strings.Builder

// 	if len(node.Arguments) > 0 {

// 		builder.WriteRune('(')

// 		for i, argument := range node.Arguments {
// 			if i > 0 {
// 				builder.WriteString(", ")
// 			}
// 			builder.WriteString(argument.Value)
// 		}

// 		builder.WriteRune(')')

// 	}

// 	return builder.String()
// }

// func (node *ast.SourceFile) String() string {
// 	var builder strings.Builder

// 	for _, node := range node.Nodes {
// 		builder.WriteString(node.String())
// 	}

// 	return builder.String()
// }

// func (node *ast.Term) String() string {
// 	var builder strings.Builder

// 	builder.WriteString(node.Left.String())

// 	builder.WriteByte(' ')

// 	switch node.Operator {
// 	case AdditiveOperatorAddition:
// 		builder.WriteString("+")
// 	case AdditiveOperatorSubtraction:
// 		builder.WriteString("-")
// 	}

// 	builder.WriteByte(' ')

// 	builder.WriteString(node.Right.String())

// 	return builder.String()
// }

// func (node *ast.Unary) String() string {
// 	var builder strings.Builder

// 	switch node.Operator {
// 	case UnaryOperatorSubtraction:
// 		builder.WriteRune('-')
// 	case UnaryOperatorNot:
// 		builder.WriteRune('!')
// 	case UnaryOperatorSpread:
// 		builder.WriteString("...")
// 	}

// 	builder.WriteString(node.Primary.String())

// 	return builder.String()
// }

// func (node *ast.VariableDeclaration) String() string {
// 	var builder strings.Builder

// 	builder.WriteString("let ")
// 	builder.WriteString(node.Identifier)

// 	if node.Expression != nil {
// 		builder.WriteString(" = ")
// 		builder.WriteString(node.Expression.String())
// 	}

// 	builder.WriteString("\n")

// 	return builder.String()
// }
