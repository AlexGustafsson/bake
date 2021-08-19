package parsing

import (
	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func parseSourceFile(parser *Parser) (*nodes.SourceFile, error) {
	parser.require(lexing.ItemStartOfInput)

	sourceFile := nodes.CreateSourceFile(nodes.Range{})

	declarations := make([]nodes.Node, 0)
dec:
	for {
		token := parser.peek()
		switch token.Type {
		case lexing.ItemNewline, lexing.ItemWhitespace:
			// Ignore
			parser.nextItem()
		case lexing.ItemEndOfInput:
			parser.nextItem()
			break dec
		case lexing.ItemKeywordPackage:
			declaration := parsePackageDeclaration(parser)
			declarations = append(declarations, declaration)
		case lexing.ItemKeywordImport:
			declaration := parseImportsDeclaration(parser)
			declarations = append(declarations, declaration)
		default:
			declaration := parseTopLevelDeclaration(parser)
			declarations = append(declarations, declaration)
		}
	}

	sourceFile.Nodes = append(sourceFile.Nodes, declarations...)

	return sourceFile, nil
}

func parsePackageDeclaration(parser *Parser) *nodes.PackageDeclaration {
	startToken := parser.require(lexing.ItemKeywordPackage)
	identifier := parser.require(lexing.ItemIdentifier)
	parser.require(lexing.ItemNewline)
	return nodes.CreatePackageDeclaration(nodes.CreateRangeFromItem(startToken), identifier.Value)
}

func parseImportsDeclaration(parser *Parser) *nodes.ImportsDeclaration {
	startToken := parser.require(lexing.ItemKeywordImport)

	parser.require(lexing.ItemLeftParentheses)
	parser.require(lexing.ItemNewline)

	imports := make([]*nodes.InterpretedString, 0)
	for {
		if token, ok := parser.expectPeek(lexing.ItemInterpretedString); ok {
			parser.nextItem()
			node := nodes.CreateInterpretedString(nodes.CreateRangeFromItem(startToken), token.Value)
			imports = append(imports, node)
			parser.require(lexing.ItemNewline)
		} else {
			break
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateImportsDeclaration(nodes.CreateRangeFromItem(startToken), imports)
}

func parseTopLevelDeclaration(parser *Parser) nodes.Node {
	token := parser.peek()

	switch token.Type {
	case lexing.ItemKeywordLet:
		return parseVariableDeclaration(parser)
	case lexing.ItemKeywordFunc:
		return parseFunctionDeclaration(parser, false)
	case lexing.ItemKeywordRule:
		parser.errorf("rule functions are not implemented")
	case lexing.ItemKeywordAlias:
		parser.errorf("aliases are not implemented")
	case lexing.ItemKeywordExport:
		parser.nextItem()
		switch token.Type {
		case lexing.ItemKeywordFunc:
			return parseFunctionDeclaration(parser, true)
		case lexing.ItemKeywordRule:
			parser.errorf("rule functions are not implemented")
		default:
			parser.tokenErrorf(token, "unexpected %s", token.Type.String())
		}
	default:
		parser.tokenErrorf(token, "unexpected %s", token.Type.String())
	}
	return nil
}

func parseVariableDeclaration(parser *Parser) nodes.Node {
	startToken := parser.nextItem()
	identifier := parser.require(lexing.ItemIdentifier)
	var expression nodes.Node = nil
	if _, ok := parser.expectPeek(lexing.ItemAssignment); ok {
		parser.nextItem()
		expression = parseExpression(parser)
	}

	return nodes.CreateVariableDeclaration(nodes.CreateRangeFromItem(startToken), identifier.Value, expression)
}

func parseFunctionDeclaration(parser *Parser, exported bool) nodes.Node {
	startToken := parser.require(lexing.ItemKeywordFunc)

	identifier := parser.require(lexing.ItemIdentifier)

	signature, _ := parseSignature(parser)

	block := parseBlock(parser)

	return nodes.CreateFunctionDeclaration(nodes.CreateRangeFromItem(startToken), exported, identifier.Value, signature, block)
}

func parseSignature(parser *Parser) (*nodes.Signature, bool) {
	if _, ok := parser.expectPeek(lexing.ItemLeftParentheses); !ok {
		return nil, false
	}

	startToken := parser.require(lexing.ItemLeftParentheses)

	arguments := make([]string, 0)
	for {
		// If there's an argument already specified, require comma separation
		if len(arguments) > 0 {
			_, ok := parser.expectPeek(lexing.ItemComma)
			if ok {
				parser.nextItem()
				token := parser.require(lexing.ItemIdentifier)
				argument := nodes.CreateIdentifier(nodes.CreateRangeFromItem(startToken), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		} else {
			token, ok := parser.expectPeek(lexing.ItemIdentifier)
			if ok {
				parser.nextItem()
				argument := nodes.CreateIdentifier(nodes.CreateRangeFromItem(startToken), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateSignature(nodes.CreateRangeFromItem(startToken), arguments), true
}

func parseBlock(parser *Parser) *nodes.Block {
	startToken := parser.require(lexing.ItemLeftCurly)

	statements := make([]nodes.Node, 0)

dec:
	for {
		token := parser.peek()
		switch token.Type {
		case lexing.ItemNewline, lexing.ItemWhitespace:
			// Ignore
			parser.nextItem()
		case lexing.ItemRightCurly:
			parser.nextItem()
			break dec
		case lexing.ItemEndOfInput:
			parser.tokenErrorf(token, "unexpected end of file, expected '}'")
		default:
			statement := parseStatement(parser)
			statements = append(statements, statement)
		}
	}

	return nodes.CreateBlock(nodes.CreateRangeFromItem(startToken), statements)
}

func parseStatement(parser *Parser) nodes.Node {
	token := parser.peek()
	if token.Type == lexing.ItemKeywordShell {
		return parseShellStatement(parser)
	} else {
		expression := parseExpression(parser)

		token := parser.peek()
		switch token.Type {
		case lexing.ItemIncrement:
			parser.nextItem()
			return nodes.CreateIncrement(nodes.CreateRange(expression.Start(), expression.End()), expression)
		case lexing.ItemDecrement:
			parser.nextItem()
			return nodes.CreateDecrement(nodes.CreateRange(expression.Start(), expression.End()), expression)
		case lexing.ItemLooseAssignment:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateLooseAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		case lexing.ItemAdditionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateAdditionAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		case lexing.ItemSubtractionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateSubtractionAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		case lexing.ItemMultiplicationAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateMultiplicationAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		case lexing.ItemDivisionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateDivisionAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		case lexing.ItemAssignment:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateAssignment(nodes.CreateRange(expression.Start(), expression.End()), expression, value)
		}

		return expression
	}
}

func parseShellStatement(parser *Parser) *nodes.ShellStatement {
	startToken := parser.require(lexing.ItemKeywordShell)

	var shellString lexing.Item
	multiline := false
	if _, ok := parser.expectPeek(lexing.ItemLeftCurly); ok {
		parser.nextItem()

		shellString = parser.require(lexing.ItemShellString)
		multiline = true

		parser.require(lexing.ItemRightCurly)
	} else {
		shellString = parser.require(lexing.ItemShellString)
	}

	return nodes.CreateShellStatement(nodes.CreateRangeFromItem(startToken), multiline, shellString.Value)
}

func parseExpression(parser *Parser) nodes.Node {
	return parseEquality(parser)
}

func parseEquality(parser *Parser) nodes.Node {
	left := parseComparison(parser)

	operatorToken := parser.peek()
	var operator nodes.EqualityOperator = -1
	switch operatorToken.Type {
	case lexing.ItemOr:
		operator = nodes.EqualityOperatorOr
	case lexing.ItemAnd:
		operator = nodes.EqualityOperatorAnd
	}

	if operator != -1 {
		parser.nextItem()
		right := parseComparison(parser)
		return nodes.CreateEquality(nodes.CreateRange(left.Start(), left.End()), operator, left, right)
	}

	return left
}

func parseComparison(parser *Parser) nodes.Node {
	left := parseTerm(parser)

	operatorToken := parser.peek()
	var operator nodes.ComparisonOperator = -1
	switch operatorToken.Type {
	case lexing.ItemEquals:
		operator = nodes.ComparisonOperatorEquals
	case lexing.ItemNotEqual:
		operator = nodes.ComparisonOperatorNotEquals
	case lexing.ItemLessThan:
		operator = nodes.ComparisonOperatorLessThan
	case lexing.ItemLessThanOrEqual:
		operator = nodes.ComparisonOperatorLessThanOrEqual
	case lexing.ItemGreaterThan:
		operator = nodes.ComparisonOperatorGreaterThan
	case lexing.ItemGreaterThanOrEqual:
		operator = nodes.ComparisonOperatorGreaterThanOrEqual
	}

	if operator != -1 {
		parser.nextItem()
		right := parseTerm(parser)
		return nodes.CreateComparison(nodes.CreateRange(left.Start(), left.End()), operator, left, right)
	}

	return left
}

func parseTerm(parser *Parser) nodes.Node {
	left := parseFactor(parser)

	operatorToken := parser.peek()
	var operator nodes.AdditiveOperator = -1
	switch operatorToken.Type {
	case lexing.ItemAddition:
		operator = nodes.AdditiveOperatorAddition
	case lexing.ItemSubtraction:
		operator = nodes.AdditiveOperatorSubtraction
	}

	if operator != -1 {
		parser.nextItem()
		right := parseFactor(parser)
		return nodes.CreateTerm(nodes.CreateRange(left.Start(), left.End()), operator, left, right)
	}

	return left
}

func parseFactor(parser *Parser) nodes.Node {
	left := parseUnary(parser)

	operatorToken := parser.peek()
	var operator nodes.MultiplicativeOperator = -1
	switch operatorToken.Type {
	case lexing.ItemMultiplication:
		operator = nodes.MultiplicativeOperatorMultiplication
	case lexing.ItemDivision:
		operator = nodes.MultiplicativeOperatorDivision
	}

	if operator != -1 {
		parser.nextItem()
		right := parseUnary(parser)
		return nodes.CreateFactor(nodes.CreateRange(left.Start(), left.End()), operator, left, right)
	}

	return left
}

func parseUnary(parser *Parser) nodes.Node {
	operatorToken := parser.peek()
	var operator nodes.UnaryOperator = -1
	switch operatorToken.Type {
	case lexing.ItemSubtraction:
		operator = nodes.UnaryOperatorSubtraction
	case lexing.ItemNot:
		operator = nodes.UnaryOperatorNot
	case lexing.ItemSpread:
		operator = nodes.UnaryOperatorSpread
	}

	if operator == -1 {
		return parsePrimary(parser)
	} else {
		parser.nextItem()
		primary := parsePrimary(parser)
		return nodes.CreateUnary(nodes.CreateRangeFromItem(operatorToken), operator, primary)
	}
}

func parsePrimary(parser *Parser) nodes.Node {
	// TODO: fix recursion
	operand := parseOperand(parser)

	token := parser.peek()
	switch token.Type {
	case lexing.ItemDot:
		startToken := parser.nextItem()
		identifier := parser.require(lexing.ItemIdentifier)
		return nodes.CreateSelector(nodes.CreateRangeFromItem(startToken), operand, identifier.Value)
	case lexing.ItemColonColon:
		startToken := parser.nextItem()
		identifier := parser.require(lexing.ItemIdentifier)
		return nodes.CreateImportSelector(nodes.CreateRangeFromItem(startToken), operand, identifier.Value)
	case lexing.ItemLeftBracket:
		startToken := parser.nextItem()
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightBracket)
		return nodes.CreateIndex(nodes.CreateRangeFromItem(startToken), operand, expression)
	case lexing.ItemLeftParentheses:
		startToken := parser.nextItem()

		arguments := make([]nodes.Node, 0)
		for {
			// If there's an argument already specified, require comma separation
			if len(arguments) > 0 {
				_, ok := parser.expectPeek(lexing.ItemComma)
				if ok {
					parser.nextItem()
					arguments = append(arguments, parseExpression(parser))
				} else {
					break
				}
			} else {
				arguments = append(arguments, parseExpression(parser))
			}
		}

		parser.require(lexing.ItemRightParentheses)

		return nodes.CreateInvokation(nodes.CreateRangeFromItem(startToken), operand, arguments)
	default:
		return operand
	}
}

func parseOperand(parser *Parser) nodes.Node {
	token := parser.nextItem()
	switch token.Type {
	case lexing.ItemInteger:
		return nodes.CreateInteger(nodes.CreateRangeFromItem(token), token.Value)
	case lexing.ItemInterpretedString:
		return nodes.CreateInterpretedString(nodes.CreateRangeFromItem(token), token.Value)
	case lexing.ItemRawString:
		return nodes.CreateRawString(nodes.CreateRangeFromItem(token), token.Value)
	case lexing.ItemIdentifier:
		return nodes.CreateIdentifier(nodes.CreateRangeFromItem(token), token.Value)
	case lexing.ItemBoolean:
		return nodes.CreateBoolean(nodes.CreateRangeFromItem(token), token.Value)
	case lexing.ItemLeftParentheses:
		// TODO: do we need to keep the parentheses?
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightParentheses)
		return expression
	case lexing.ItemLeftBracket:
		array := nodes.CreateArray(nodes.CreateRangeFromItem(token), make([]nodes.Node, 0))
	dec:
		for {
			token := parser.peek()
			switch token.Type {
			case lexing.ItemNewline:
				// Ignore
				parser.nextItem()
			case lexing.ItemEndOfInput:
				parser.tokenErrorf(token, "unexpected end of input - missing ']'")
			case lexing.ItemRightBracket:
				parser.nextItem()
				break dec
			case lexing.ItemComma:
				parser.nextItem()
				if len(array.Elements) == 0 {
					parser.tokenErrorf(token, "unexpected comma - missing expression")
				}

				expression := parseExpression(parser)
				array.Elements = append(array.Elements, expression)
			default:
				expression := parseExpression(parser)
				array.Elements = append(array.Elements, expression)
			}
		}

		return array
	default:
		parser.tokenErrorf(token, "expected operand, got '%s'", token.Type.String())
		return nil
	}
}
