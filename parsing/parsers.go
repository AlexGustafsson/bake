package parsing

import (
	"github.com/AlexGustafsson/bake/lexing"
)

func parseRoot(parser *Parser) stateModifier {
	switch item := parser.lexer.NextItem(); item.Type {
	case lexing.ItemWhitespace, lexing.ItemNewline:
		// Skip whitespace
		return parseRoot
	case lexing.ItemImport:
		return parseImport
	case lexing.ItemEndOfFile:
		return nil
	case lexing.ItemError:
		parser.Errorf("Got error token: %s\n", item)
		return nil
	default:
		parser.Errorf("Unexpected token: %s\n", item.String())
		return nil
	}
}

func parseImport(parser *Parser) stateModifier {
	if item := parser.lexer.NextNonWhitespaceItem(true); item.Type != lexing.ItemLeftParentheses {
		parser.Errorf("Expected left parentheses\n")
		return nil
	}

	for {
		item := parser.lexer.NextNonWhitespaceItem(true)
		if item.Type == lexing.ItemString {
			parser.syntaxTree.Imports = append(parser.syntaxTree.Imports, item.Value)
		} else if item.Type == lexing.ItemRightParentheses {
			break
		} else {
			parser.Errorf("Expected right parentheses or string, got %s - %s\n", item.Type.String(), item.Message)
			return nil
		}
	}

	return parseRoot
}
