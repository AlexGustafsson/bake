package parsing

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/lexing"
)

func parseSourceFile(parser *Parser) *ast.Block {
	parser.require(lexing.ItemStartOfInput)

	declarations := make([]ast.Node, 0)
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

	return ast.CreateBlock(ast.CreateRange(ast.Position{0, 0, 0}, ast.Position{0, 0, 0}), declarations)
}

func parsePackageDeclaration(parser *Parser) *ast.PackageDeclaration {
	startToken := parser.require(lexing.ItemKeywordPackage)
	identifier := parser.require(lexing.ItemIdentifier)
	parser.require(lexing.ItemNewline)
	return ast.CreatePackageDeclaration(createRangeFromItem(startToken), identifier.Value)
}

func parseImportsDeclaration(parser *Parser) *ast.ImportsDeclaration {
	startToken := parser.require(lexing.ItemKeywordImport)

	parser.require(lexing.ItemLeftParentheses)
	parser.require(lexing.ItemNewline)

	imports := make([]*ast.EvaluatedString, 0)
	for {
		if _, ok := parser.expectPeek(lexing.ItemDoubleQuote); ok {
			node := parseEvaluatedString(parser)
			imports = append(imports, node)
			parser.require(lexing.ItemNewline)
		} else {
			break
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return ast.CreateImportsDeclaration(createRangeFromItem(startToken), imports)
}

func parseTopLevelDeclaration(parser *Parser) ast.Node {
	token := parser.peek()

	switch token.Type {
	case lexing.ItemKeywordRule:
		return parseRuleFuncionDeclaration(parser, false)
	case lexing.ItemKeywordAlias:
		return parseAliasDeclaration(parser, false)
	case lexing.ItemKeywordExport:
		parser.nextItem()
		token := parser.peek()
		switch token.Type {
		case lexing.ItemKeywordFunc:
			return parseFunctionDeclaration(parser, true)
		case lexing.ItemKeywordRule:
			return parseRuleFuncionDeclaration(parser, true)
		case lexing.ItemKeywordAlias:
			return parseAliasDeclaration(parser, true)
		default:
			parser.tokenErrorf(token, "unexpected %s", token.Type.String())
		}
	case lexing.ItemDoubleQuote, lexing.ItemRawString, lexing.ItemLeftBracket:
		return parseRule(parser)
	default:
		return parseStatement(parser)
	}
	return nil
}

func parseVariableDeclaration(parser *Parser) ast.Node {
	startToken := parser.nextItem()
	identifier := parser.require(lexing.ItemIdentifier)
	var expression ast.Node = nil
	if _, ok := parser.expectPeek(lexing.ItemAssignment); ok {
		parser.nextItem()
		expression = parseExpression(parser)
	}

	return ast.CreateVariableDeclaration(createRangeFromItem(startToken), identifier.Value, expression)
}

func parseFunctionDeclaration(parser *Parser, exported bool) ast.Node {
	startToken := parser.require(lexing.ItemKeywordFunc)

	identifier := parser.require(lexing.ItemIdentifier)

	signature, _ := parseSignature(parser)

	block := parseBlock(parser)

	return ast.CreateFunctionDeclaration(createRangeFromItem(startToken), exported, identifier.Value, signature, block)
}

func parseRuleFuncionDeclaration(parser *Parser, exported bool) ast.Node {
	startToken := parser.require(lexing.ItemKeywordRule)

	identifier := parser.require(lexing.ItemIdentifier)

	signature, _ := parseSignature(parser)

	block := parseBlock(parser)

	return ast.CreateRuleFunctionDeclaration(createRangeFromItem(startToken), exported, identifier.Value, signature, block)
}

func parseAliasDeclaration(parser *Parser, exported bool) ast.Node {
	startToken := parser.require(lexing.ItemKeywordAlias)

	identifier := parser.nextItem()
	// Allow keywords not lexed as identifiers to be used in selections
	if identifier.Type != lexing.ItemIdentifier && !identifier.IsKeyword() {
		parser.tokenErrorf(identifier, "expected identifier, got '%s'", identifier.Value)
	}

	parser.require(lexing.ItemColon)

	expression := parseExpression(parser)

	return ast.CreateAliasDeclaration(createRangeFromItem(startToken), exported, identifier.String(), expression)
}

func parseSignature(parser *Parser) (*ast.Signature, bool) {
	if _, ok := parser.expectPeek(lexing.ItemLeftParentheses); !ok {
		return nil, false
	}

	startToken := parser.require(lexing.ItemLeftParentheses)

	arguments := make([]*ast.Identifier, 0)
	for {
		// If there's an argument already specified, require comma separation
		if len(arguments) > 0 {
			_, ok := parser.expectPeek(lexing.ItemComma)
			if ok {
				parser.nextItem()
				token := parser.require(lexing.ItemIdentifier)
				argument := ast.CreateIdentifier(createRangeFromItem(startToken), token.Value)
				arguments = append(arguments, argument)
			} else {
				break
			}
		} else {
			token, ok := parser.expectPeek(lexing.ItemIdentifier)
			if ok {
				parser.nextItem()
				argument := ast.CreateIdentifier(createRangeFromItem(startToken), token.Value)
				arguments = append(arguments, argument)
			} else {
				break
			}
		}
	}

	parser.require(lexing.ItemRightParentheses)

	return ast.CreateSignature(createRangeFromItem(startToken), arguments), true
}

func parseBlock(parser *Parser) *ast.Block {
	startToken := parser.require(lexing.ItemLeftCurly)

	statements := make([]ast.Node, 0)

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

	return ast.CreateBlock(createRangeFromItem(startToken), statements)
}

func parseRule(parser *Parser) *ast.RuleDeclaration {
	outputs := make([]ast.Node, 0)
	dependencies := make([]ast.Node, 0)
	var block *ast.Block = nil
	var derived ast.Node = nil

	// Parse a string literal or an array as the outputs
	startToken := parser.peek()
	switch startToken.Type {
	case lexing.ItemDoubleQuote:
		node := parseEvaluatedString(parser)
		outputs = append(outputs, node)
	case lexing.ItemRawString:
		parser.nextItem()
		outputs = append(outputs, ast.CreateRawString(createRangeFromItem(startToken), startToken.Value))
	case lexing.ItemLeftBracket:
		array := parseArray(parser)
		outputs = array.Elements
	}

	// Optionally parse an array of dependencies
	if _, ok := parser.expectPeek(lexing.ItemLeftBracket); ok {
		array := parseArray(parser)
		dependencies = array.Elements
	}

	// Parse either a rule function usage with an optional block, or a required block
	if _, ok := parser.expectPeek(lexing.ItemColon); ok {
		parser.require(lexing.ItemColon)
		derived = parsePrimary(parser)

		// TODO: May need to take newlines etc. into account?
		if _, ok := parser.expectPeek(lexing.ItemLeftCurly); ok {
			block = parseBlock(parser)
		}
	} else {
		block = parseBlock(parser)
	}

	return ast.CreateRuleDeclaration(createRangeFromItem(startToken), outputs, dependencies, derived, block)
}

func parseStatement(parser *Parser) ast.Node {
	token := parser.peek()
	switch token.Type {
	case lexing.ItemKeywordFunc:
		return parseFunctionDeclaration(parser, false)
	case lexing.ItemKeywordLet:
		return parseVariableDeclaration(parser)
	case lexing.ItemKeywordShell:
		return parseShellStatement(parser)
	case lexing.ItemKeywordReturn:
		startToken := parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateReturnStatement(createRangeFromItem(startToken), value)
	case lexing.ItemKeywordIf:
		return parseIfStatement(parser)
	case lexing.ItemKeywordFor:
		return parseForStatement(parser)
	default:
		return parseSimpleStatement(parser)
	}
}

func parseIfStatement(parser *Parser) *ast.IfStatement {
	startToken := parser.require(lexing.ItemKeywordIf)
	expression := parseExpression(parser)
	positiveBranch := parseBlock(parser)

	var negativeBranch ast.Node = nil
	if _, ok := parser.expectPeek(lexing.ItemKeywordElse); ok {
		parser.nextItem()
		if _, ok := parser.expectPeek(lexing.ItemKeywordIf); ok {
			negativeBranch = parseIfStatement(parser)
		} else {
			negativeBranch = parseBlock(parser)
		}
	}

	return ast.CreateIfStatement(createRangeFromItem(startToken), expression, positiveBranch, negativeBranch)
}

func parseForStatement(parser *Parser) *ast.ForStatement {
	startToken := parser.require(lexing.ItemKeywordFor)
	identifier := parser.require(lexing.ItemIdentifier)
	parser.require(lexing.ItemKeywordIn)
	expression := parseExpression(parser)
	block := parseBlock(parser)

	return ast.CreateForStatement(createRangeFromItem(startToken), ast.CreateIdentifier(createRangeFromItem(identifier), identifier.Value), expression, block)
}

func parseSimpleStatement(parser *Parser) ast.Node {
	expression := parseExpression(parser)

	token := parser.peek()
	switch token.Type {
	case lexing.ItemIncrement:
		parser.nextItem()
		return ast.CreateIncrement(expression.Range(), expression)
	case lexing.ItemDecrement:
		parser.nextItem()
		return ast.CreateDecrement(expression.Range(), expression)
	case lexing.ItemLooseAssignment:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateLooseAssignment(expression.Range(), expression, value)
	case lexing.ItemAdditionAssign:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateAdditionAssignment(expression.Range(), expression, value)
	case lexing.ItemSubtractionAssign:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateSubtractionAssignment(expression.Range(), expression, value)
	case lexing.ItemMultiplicationAssign:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateMultiplicationAssignment(expression.Range(), expression, value)
	case lexing.ItemDivisionAssign:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateDivisionAssignment(expression.Range(), expression, value)
	case lexing.ItemModuloAssign:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateModuloAssignment(expression.Range(), expression, value)
	case lexing.ItemAssignment:
		parser.nextItem()
		value := parseExpression(parser)
		return ast.CreateAssignment(expression.Range(), expression, value)
	}

	return expression
}

func parseShellStatement(parser *Parser) *ast.ShellStatement {
	startToken := parser.require(lexing.ItemKeywordShell)
	multiline := false
	parts := make([]ast.Node, 0)

	if _, ok := parser.expectPeek(lexing.ItemLeftCurly); ok {
		multiline = true
		parser.nextItem()
	}

	for {
		item := parser.nextItem()
		switch item.Type {
		case lexing.ItemStringPart:
			parts = append(parts, ast.CreateStringPart(createRangeFromItem(item), item.Value))
		case lexing.ItemSubstitutionStart:
			expression := parseExpression(parser)
			parser.require(lexing.ItemSubstitutionEnd)
			parts = append(parts, expression)
		case lexing.ItemNewline:
			return ast.CreateShellStatement(createRangeFromItem(startToken), multiline, parts)
		case lexing.ItemRightCurly:
			endToken := item
			start := createRangeFromItem(startToken)
			end := createRangeFromItem(endToken)
			r := ast.CreateRange(start.Start, end.End)
			return ast.CreateShellStatement(r, multiline, parts)
		}
	}
}

func parseExpression(parser *Parser) ast.Node {
	return parseEquality(parser)
}

func parseEquality(parser *Parser) ast.Node {
	left := parseComparison(parser)

	for {
		operatorToken := parser.peek()
		var operator ast.EqualityOperator = -1
		switch operatorToken.Type {
		case lexing.ItemOr:
			operator = ast.EqualityOperatorOr
		case lexing.ItemAnd:
			operator = ast.EqualityOperatorAnd
		}

		if operator == -1 {
			break
		} else {
			parser.nextItem()
			right := parseComparison(parser)
			left = ast.CreateEquality(left.Range(), operator, left, right)
		}
	}

	return left
}

func parseComparison(parser *Parser) ast.Node {
	left := parseTerm(parser)

	for {
		operatorToken := parser.peek()
		var operator ast.ComparisonOperator = -1
		switch operatorToken.Type {
		case lexing.ItemEquals:
			operator = ast.ComparisonOperatorEquals
		case lexing.ItemNotEqual:
			operator = ast.ComparisonOperatorNotEquals
		case lexing.ItemLessThan:
			operator = ast.ComparisonOperatorLessThan
		case lexing.ItemLessThanOrEqual:
			operator = ast.ComparisonOperatorLessThanOrEqual
		case lexing.ItemGreaterThan:
			operator = ast.ComparisonOperatorGreaterThan
		case lexing.ItemGreaterThanOrEqual:
			operator = ast.ComparisonOperatorGreaterThanOrEqual
		}

		if operator == -1 {
			break
		} else {
			parser.nextItem()
			right := parseTerm(parser)
			left = ast.CreateComparison(left.Range(), operator, left, right)
		}

	}

	return left
}

func parseTerm(parser *Parser) ast.Node {
	left := parseFactor(parser)

	for {
		operatorToken := parser.peek()
		var operator ast.AdditiveOperator = -1
		switch operatorToken.Type {
		case lexing.ItemAddition:
			operator = ast.AdditiveOperatorAddition
		case lexing.ItemSubtraction:
			operator = ast.AdditiveOperatorSubtraction
		}

		if operator == -1 {
			break
		} else {
			parser.nextItem()
			right := parseFactor(parser)
			left = ast.CreateTerm(left.Range(), operator, left, right)
		}
	}

	return left
}

func parseFactor(parser *Parser) ast.Node {
	left := parseUnary(parser)

	for {
		operatorToken := parser.peek()
		var operator ast.MultiplicativeOperator = -1
		switch operatorToken.Type {
		case lexing.ItemMultiplication:
			operator = ast.MultiplicativeOperatorMultiplication
		case lexing.ItemDivision:
			operator = ast.MultiplicativeOperatorDivision
		case lexing.ItemModulo:
			operator = ast.MultiplicativeOperatorModulo
		}

		if operator == -1 {
			break
		} else {
			parser.nextItem()
			right := parseUnary(parser)
			left = ast.CreateFactor(left.Range(), operator, left, right)
		}
	}

	return left
}

func parseUnary(parser *Parser) ast.Node {
	operatorToken := parser.peek()
	var operator ast.UnaryOperator = -1
	switch operatorToken.Type {
	case lexing.ItemSubtraction:
		operator = ast.UnaryOperatorSubtraction
	case lexing.ItemNot:
		operator = ast.UnaryOperatorNot
	case lexing.ItemSpread:
		operator = ast.UnaryOperatorSpread
	}

	if operator == -1 {
		return parsePrimary(parser)
	} else {
		parser.nextItem()
		primary := parsePrimary(parser)
		return ast.CreateUnary(createRangeFromItem(operatorToken), operator, primary)
	}
}

func parsePrimary(parser *Parser) ast.Node {
	left := parseOperand(parser)

dec:
	for {
		token := parser.peek()
		switch token.Type {
		case lexing.ItemDot:
			startToken := parser.nextItem()
			identifier := parser.nextItem()
			// Allow keywords not lexed as identifiers to be used in selections
			if identifier.Type != lexing.ItemIdentifier && !identifier.IsKeyword() {
				parser.tokenErrorf(identifier, "expected identifier, got '%s'", identifier.Value)
			}
			left = ast.CreateSelector(createRangeFromItem(startToken), left, identifier.Value)
		case lexing.ItemLeftBracket:
			startToken := parser.nextItem()
			expression := parseExpression(parser)
			parser.require(lexing.ItemRightBracket)
			left = ast.CreateIndex(createRangeFromItem(startToken), left, expression)
		case lexing.ItemLeftParentheses:
			arguments := parseExpressionList(parser, lexing.ItemLeftParentheses, lexing.ItemRightParentheses)
			left = ast.CreateInvocation(createRangeFromItem(token), left, arguments)
		default:
			break dec
		}
	}

	return left
}

func parseOperand(parser *Parser) ast.Node {
	token := parser.peek()
	switch token.Type {
	case lexing.ItemInteger:
		parser.nextItem()
		return ast.CreateInteger(createRangeFromItem(token), token.Value)
	case lexing.ItemDoubleQuote:
		return parseEvaluatedString(parser)
	case lexing.ItemRawString:
		parser.nextItem()
		return ast.CreateRawString(createRangeFromItem(token), token.Value)
	case lexing.ItemIdentifier:
		parser.nextItem()
		if _, ok := parser.expectPeek(lexing.ItemColonColon); ok {
			parser.nextItem()
			identifier := parser.require(lexing.ItemIdentifier)
			return ast.CreateImportSelector(createRangeFromItem(token), token.Value, identifier.Value)
		} else {
			return ast.CreateIdentifier(createRangeFromItem(token), token.Value)
		}
	case lexing.ItemBoolean:
		parser.nextItem()
		return ast.CreateBoolean(createRangeFromItem(token), token.Value)
	case lexing.ItemLeftParentheses:
		parser.nextItem()
		// TODO: do we need to keep the parentheses?
		expression := parseExpression(parser)
		parser.require(lexing.ItemRightParentheses)
		return expression
	case lexing.ItemLeftBracket:
		return parseArray(parser)
	default:
		// TODO: do we need to consume the peeked token here?
		parser.tokenErrorf(token, "expected operand, got '%s'", token.Type.String())
		return nil
	}
}

func parseArray(parser *Parser) *ast.Array {
	startToken := parser.peek()
	array := ast.CreateArray(createRangeFromItem(startToken), make([]ast.Node, 0))
	array.Elements = parseExpressionList(parser, lexing.ItemLeftBracket, lexing.ItemRightBracket)
	return array
}

func parseExpressionList(parser *Parser, start lexing.ItemType, end lexing.ItemType) []ast.Node {
	parser.require(start)
	expressions := make([]ast.Node, 0)
dec:
	for {
		token := parser.peek()
		switch token.Type {
		case lexing.ItemNewline:
			// Ignore
			parser.nextItem()
		case lexing.ItemEndOfInput:
			parser.tokenErrorf(token, "unexpected end of input - missing '%s'", end)
		case end:
			parser.nextItem()
			break dec
		case lexing.ItemComma:
			parser.nextItem()
			if len(expressions) == 0 {
				parser.tokenErrorf(token, "unexpected comma - missing expression")
			}

			// Don't allow two commas after each other - this is done here so that the main switch
			// loop may be reused each iteration to handle whitespace uniformly etc.
			if token, ok := parser.expectPeek(lexing.ItemComma); ok {
				parser.tokenErrorf(token, "unexpected comma")
			}
		default:
			expression := parseExpression(parser)
			expressions = append(expressions, expression)
		}
	}

	return expressions
}

func parseEvaluatedString(parser *Parser) *ast.EvaluatedString {
	startToken := parser.require(lexing.ItemDoubleQuote)

	children := make([]ast.Node, 0)

	for {
		item := parser.nextItem()
		switch item.Type {
		case lexing.ItemStringPart:
			children = append(children, ast.CreateStringPart(createRangeFromItem(item), item.Value))
		case lexing.ItemSubstitutionStart:
			expression := parseExpression(parser)
			parser.require(lexing.ItemSubstitutionEnd)
			children = append(children, expression)
		case lexing.ItemDoubleQuote:
			endToken := item
			start := createRangeFromItem(startToken)
			end := createRangeFromItem(endToken)
			r := ast.CreateRange(start.Start, end.End)
			return ast.CreateEvaluatedString(r, children)
		}
	}
}
