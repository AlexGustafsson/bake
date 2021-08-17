package parsing

import (
	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func parseSourceFile(parser *Parser) (*nodes.SourceFile, error) {
	parser.require(lexing.ItemStartOfInput)

	sourceFile := nodes.CreateSourceFile(0)

	if packageDeclaration, ok := parsePackageDeclaration(parser); ok {
		sourceFile.Nodes = append(sourceFile.Nodes, packageDeclaration)
	}

	if importsDeclaration, ok := parseImportsDeclaration(parser); ok {
		sourceFile.Nodes = append(sourceFile.Nodes, importsDeclaration)
	}

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
		default:
			declaration := parseTopLevelDeclaration(parser)
			declarations = append(declarations, declaration)
		}
	}

	sourceFile.Nodes = append(sourceFile.Nodes, declarations...)

	return sourceFile, nil
}

func parsePackageDeclaration(parser *Parser) (*nodes.PackageDeclaration, bool) {
	if _, ok := parser.expectPeek(lexing.ItemKeywordPackage); !ok {
		return nil, false
	}

	startToken := parser.require(lexing.ItemKeywordPackage)
	identifier := parser.require(lexing.ItemIdentifier)
	parser.require(lexing.ItemNewline)
	return nodes.CreatePackageDeclaration(nodes.NodePosition(startToken.Start), identifier), true
}

func parseImportsDeclaration(parser *Parser) (*nodes.ImportsDeclaration, bool) {
	if _, ok := parser.expectPeek(lexing.ItemKeywordImport); !ok {
		return nil, false
	}

	startToken := parser.require(lexing.ItemKeywordImport)

	parser.require(lexing.ItemLeftParentheses)
	parser.require(lexing.ItemNewline)

	imports := make([]*nodes.InterpretedString, 0)
	for {
		if token, ok := parser.expectPeek(lexing.ItemInterpretedString); ok {
			parser.nextItem()
			node := nodes.CreateInterpretedString(nodes.NodePosition(startToken.Start), token.Value)
			imports = append(imports, node)
			parser.require(lexing.ItemNewline)
		} else {
			break
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateImportsDeclaration(nodes.NodePosition(startToken.Start), imports), true
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
			parser.errorf("unexpected token '%s'", token.String())
		}
	default:
		parser.errorf("unexpected token '%s'", token.String())
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

	return nodes.CreateVariableDeclaration(nodes.NodePosition(startToken.Start), identifier.Value, expression)
}

func parseFunctionDeclaration(parser *Parser, exported bool) nodes.Node {
	startToken := parser.require(lexing.ItemKeywordFunc)

	identifier := parser.require(lexing.ItemIdentifier)

	signature, _ := parseSignature(parser)

	block := parseBlock(parser)

	return nodes.CreateFunctionDeclaration(nodes.NodePosition(startToken.Start), exported, identifier.Value, signature, block)
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
				argument := nodes.CreateIdentifier(nodes.NodePosition(token.Start), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		} else {
			token, ok := parser.expectPeek(lexing.ItemIdentifier)
			if ok {
				parser.nextItem()
				argument := nodes.CreateIdentifier(nodes.NodePosition(token.Start), token.Value)
				arguments = append(arguments, argument.Value)
			} else {
				break
			}
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return nodes.CreateSignature(nodes.NodePosition(startToken.Start), arguments), true
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
			parser.errorf("unexpected end of file, expected '}'")
		default:
			statement := parseStatement(parser)
			statements = append(statements, statement)
		}
	}

	return nodes.CreateBlock(nodes.NodePosition(startToken.Start), statements)
}

func parseStatement(parser *Parser) nodes.Node {
	return parseExpression(parser)
}

func parseExpression(parser *Parser) nodes.Node {
	// TODO: operator precedence and recursion
	startToken := parser.peek()
	left := parseUnaryExpression(parser)

	operatorToken := parser.peek()
	var operator nodes.BinaryOperator
	switch operatorToken.Type {
	case lexing.ItemOr:
		parser.nextItem()
		operator = nodes.BinaryOperatorOr
	case lexing.ItemAnd:
		parser.nextItem()
		operator = nodes.BinaryOperatorAnd
	case lexing.ItemEquals:
		parser.nextItem()
		operator = nodes.BinaryOperatorEquals
	case lexing.ItemNotEqual:
		parser.nextItem()
		operator = nodes.BinaryOperatorNotEquals
	case lexing.ItemLessThan:
		parser.nextItem()
		operator = nodes.BinaryOperatorLessThan
	case lexing.ItemLessThanOrEqual:
		parser.nextItem()
		operator = nodes.BinaryOperatorLessThanOrEqual
	case lexing.ItemGreaterThan:
		parser.nextItem()
		operator = nodes.BinaryOperatorGreaterThan
	case lexing.ItemGreaterThanOrEqual:
		parser.nextItem()
		operator = nodes.BinaryOperatorGreaterThanOrEqual
	case lexing.ItemAddition:
		parser.nextItem()
		operator = nodes.BinaryOperatorAddition
	case lexing.ItemSubtraction:
		parser.nextItem()
		operator = nodes.BinaryOperatorSubtraction
	case lexing.ItemMultiplication:
		parser.nextItem()
		operator = nodes.BinaryOperatorMultiplication
	case lexing.ItemDivision:
		parser.nextItem()
		operator = nodes.BinaryOperatorDivision
	default:
		// If no valid operator was found, assume it was a unary expression
		return left
	}

	right := parseExpression(parser)

	return nodes.CreateBinaryExpression(nodes.NodePosition(startToken.Start), operator, left, right)
}

func parseUnaryExpression(parser *Parser) nodes.Node {
	token := parser.peek()
	switch token.Type {
	case lexing.ItemSubtraction:
		parser.nextItem()
		primary := parsePrimaryExpression(parser)
		return nodes.CreateUnaryExpression(nodes.NodePosition(token.Start), nodes.UnaryOperatorSubtraction, primary)
	case lexing.ItemNot:
		parser.nextItem()
		primary := parsePrimaryExpression(parser)
		return nodes.CreateUnaryExpression(nodes.NodePosition(token.Start), nodes.UnaryOperatorNot, primary)
	case lexing.ItemSpread:
		parser.nextItem()
		primary := parsePrimaryExpression(parser)
		return nodes.CreateUnaryExpression(nodes.NodePosition(token.Start), nodes.UnaryOperatorSpread, primary)
	default:
		return parsePrimaryExpression(parser)
	}
}

func parsePrimaryExpression(parser *Parser) nodes.Node {
	// TODO: fix recursion
	operand := parseOperand(parser)

	token := parser.peek()
	switch token.Type {
	case lexing.ItemDot:
		startToken := parser.nextItem()
		identifier := parser.require(lexing.ItemIdentifier)
		return nodes.CreateSelector(nodes.NodePosition(startToken.Start), operand, identifier.Value)
	case lexing.ItemColonColon:
		startToken := parser.nextItem()
		identifier := parser.require(lexing.ItemIdentifier)
		return nodes.CreateImportSelector(nodes.NodePosition(startToken.Start), operand, identifier.Value)
	case lexing.ItemLeftBracket:
		startToken := parser.nextItem()
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightBracket)
		return nodes.CreateIndex(nodes.NodePosition(startToken.Start), operand, expression)
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

		return nodes.CreateInvokation(nodes.NodePosition(startToken.Start), operand, arguments)
	default:
		return operand
	}
}

func parseOperand(parser *Parser) nodes.Node {
	token := parser.nextItem()
	switch token.Type {
	case lexing.ItemInteger:
		return nodes.CreateInteger(nodes.NodePosition(token.Start), token.Value)
	case lexing.ItemInterpretedString:
		return nodes.CreateInterpretedString(nodes.NodePosition(token.Start), token.Value)
	case lexing.ItemRawString:
		return nodes.CreateRawString(nodes.NodePosition(token.Start), token.Value)
	case lexing.ItemIdentifier:
		return nodes.CreateIdentifier(nodes.NodePosition(token.Start), token.Value)
	case lexing.ItemBoolean:
		return nodes.CreateBoolean(nodes.NodePosition(token.Start), token.Value)
	case lexing.ItemLeftParentheses:
		// TODO: do we need to keep the parentheses?
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightParentheses)
		return expression
	default:
		parser.errorf("expected operand, found '%s'", token.Type.String())
		return nil
	}
}
