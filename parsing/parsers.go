package parsing

import (
	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func nodePosition(item lexing.Item) nodes.NodePosition {
	return nodes.CreateNodePosition(item.Start, item.Line, item.Column)
}

func parseSourceFile(parser *Parser) (*nodes.SourceFile, error) {
	parser.require(lexing.ItemStartOfInput)

	sourceFile := nodes.CreateSourceFile(nodes.CreateNodePosition(0, 0, 0))

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
	return nodes.CreatePackageDeclaration(nodePosition(startToken), identifier.Value)
}

func parseImportsDeclaration(parser *Parser) *nodes.ImportsDeclaration {
	startToken := parser.require(lexing.ItemKeywordImport)

	parser.require(lexing.ItemLeftParentheses)
	parser.require(lexing.ItemNewline)

	imports := make([]*nodes.InterpretedString, 0)
	for {
		if token, ok := parser.expectPeek(lexing.ItemInterpretedString); ok {
			parser.nextItem()
			node := nodes.CreateInterpretedString(nodePosition(startToken), token.Value)
			imports = append(imports, node)
			parser.require(lexing.ItemNewline)
		} else {
			break
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateImportsDeclaration(nodePosition(startToken), imports)
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

	return nodes.CreateVariableDeclaration(nodePosition(startToken), identifier.Value, expression)
}

func parseFunctionDeclaration(parser *Parser, exported bool) nodes.Node {
	startToken := parser.require(lexing.ItemKeywordFunc)

	identifier := parser.require(lexing.ItemIdentifier)

	signature, _ := parseSignature(parser)

	block := parseBlock(parser)

	return nodes.CreateFunctionDeclaration(nodePosition(startToken), exported, identifier.Value, signature, block)
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
				argument := nodes.CreateIdentifier(nodePosition(startToken), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		} else {
			token, ok := parser.expectPeek(lexing.ItemIdentifier)
			if ok {
				parser.nextItem()
				argument := nodes.CreateIdentifier(nodePosition(startToken), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateSignature(nodePosition(startToken), arguments), true
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

	return nodes.CreateBlock(nodePosition(startToken), statements)
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
			return nodes.CreateIncrement(expression.Position(), expression)
		case lexing.ItemDecrement:
			parser.nextItem()
			return nodes.CreateDecrement(expression.Position(), expression)
		case lexing.ItemLooseAssignment:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateLooseAssignment(expression.Position(), expression, value)
		case lexing.ItemAdditionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateAdditionAssignment(expression.Position(), expression, value)
		case lexing.ItemSubtractionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateSubtractionAssignment(expression.Position(), expression, value)
		case lexing.ItemMultiplicationAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateMultiplicationAssignment(expression.Position(), expression, value)
		case lexing.ItemDivisionAssign:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateDivisionAssignment(expression.Position(), expression, value)
		case lexing.ItemAssignment:
			parser.nextItem()
			value := parseExpression(parser)
			return nodes.CreateAssignment(expression.Position(), expression, value)
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

	return nodes.CreateShellStatement(nodePosition(startToken), multiline, shellString.Value)
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
		return nodes.CreateEquality(left.Position(), operator, left, right)
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
		return nodes.CreateComparison(left.Position(), operator, left, right)
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
		return nodes.CreateTerm(left.Position(), operator, left, right)
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
		return nodes.CreateFactor(left.Position(), operator, left, right)
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
		return nodes.CreateUnary(nodePosition(operatorToken), operator, primary)
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
		return nodes.CreateSelector(nodePosition(startToken), operand, identifier.Value)
	case lexing.ItemColonColon:
		startToken := parser.nextItem()
		identifier := parser.require(lexing.ItemIdentifier)
		return nodes.CreateImportSelector(nodePosition(startToken), operand, identifier.Value)
	case lexing.ItemLeftBracket:
		startToken := parser.nextItem()
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightBracket)
		return nodes.CreateIndex(nodePosition(startToken), operand, expression)
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

		return nodes.CreateInvokation(nodePosition(startToken), operand, arguments)
	default:
		return operand
	}
}

func parseOperand(parser *Parser) nodes.Node {
	token := parser.nextItem()
	switch token.Type {
	case lexing.ItemInteger:
		return nodes.CreateInteger(nodePosition(token), token.Value)
	case lexing.ItemInterpretedString:
		return nodes.CreateInterpretedString(nodePosition(token), token.Value)
	case lexing.ItemRawString:
		return nodes.CreateRawString(nodePosition(token), token.Value)
	case lexing.ItemIdentifier:
		return nodes.CreateIdentifier(nodePosition(token), token.Value)
	case lexing.ItemBoolean:
		return nodes.CreateBoolean(nodePosition(token), token.Value)
	case lexing.ItemLeftParentheses:
		// TODO: do we need to keep the parentheses?
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightParentheses)
		return expression
	default:
		parser.tokenErrorf(token, "expected operand, got '%s'", token.Type.String())
		return nil
	}
}
