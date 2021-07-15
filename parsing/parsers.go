package parsing

import (
	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

func parseRoot(parser *Parser) (nodes.Node, error) {
	switch item := parser.lexer.NextNonWhitespaceItem(true); item.Type {
	case lexing.ItemImport:
		return parseImport(parser)
	case lexing.ItemEndOfFile:
		return nil, nil
	case lexing.ItemError:
		parser.errorf("Got error token: %s\n", item)
	default:
		parser.errorf("Unexpected token: %s\n", item.String())
	}

	return nil, nil
}

func parseImport(parser *Parser) (nodes.Node, error) {
	node := &nodes.NodeImport{
		NodeType: nodes.NodeImportType,
		Imports:  make([]string, 0),
	}

	if item := parser.lexer.NextNonWhitespaceItem(true); item.Type != lexing.ItemLeftParentheses {
		parser.errorf("Expected left parentheses\n")
	}

	for {
		item := parser.lexer.NextNonWhitespaceItem(true)
		if item.Type == lexing.ItemString {
			node.Imports = append(node.Imports, item.Value)
		} else if item.Type == lexing.ItemRightParentheses {
			break
		} else {
			parser.errorf("Expected right parentheses or string, got %s - %s\n", item.Type.String(), item.Message)
		}
	}

	return node, nil
}
