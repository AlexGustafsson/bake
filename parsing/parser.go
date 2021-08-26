package parsing

import (
	"container/list"
	"fmt"
	"runtime"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/lexing"
)

type Parser struct {
	lexer  *lexing.Lexer
	peeked *list.List
	input  string
}

func CreateParser() *Parser {
	return &Parser{
		lexer:  nil,
		peeked: list.New(),
	}
}

func Parse(input string) (*ast.Block, error) {
	parser := CreateParser()
	parser.input = input
	return parser.Parse(input)
}

func (parser *Parser) Parse(input string) (_ *ast.Block, err error) {
	parser.lexer = lexing.Lex(input)

	// Recover from panics caused when parsing
	defer parser.recover(&err)

	return parseSourceFile(parser), nil
}

func (parser *Parser) errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func (parser *Parser) tokenErrorf(item lexing.Item, format string, args ...interface{}) {
	r := createRangeFromItem(item)
	panic(ast.CreateTreeError(r, format, args...))
}

func (parser *Parser) recover(errp *error) {
	err := recover()
	if err != nil {
		if _, ok := err.(runtime.Error); ok {
			panic(err)
		}

		if parser != nil {
			*errp = err.(error)
		}
	}
}

func (parser *Parser) require(itemType lexing.ItemType) lexing.Item {
	item, ok := parser.expect(itemType)
	if !ok {
		parser.tokenErrorf(item, "expected %s got %s", itemType.String(), item.Type.String())
	}
	return item
}

func (parser *Parser) expect(itemType lexing.ItemType) (_ lexing.Item, ok bool) {
	item := parser.nextItem()

	if item.Type == lexing.ItemError {
		parser.tokenErrorf(item, item.Message)
	}

	if item.Type != itemType {
		return item, false
	}
	return item, true
}

func (parser *Parser) peek() lexing.Item {
	item := parser.nextItem()
	parser.peeked.PushBack(item)
	return parser.peeked.Front().Value.(lexing.Item)
}

func (parser *Parser) expectPeek(itemType lexing.ItemType) (lexing.Item, bool) {
	item := parser.peek()
	if item.Type != itemType {
		return item, false
	}
	return item, true
}

func (parser *Parser) nextItem() lexing.Item {
	if parser.peeked.Len() > 0 {
		element := parser.peeked.Front()
		return parser.peeked.Remove(element).(lexing.Item)
	}

	// TODO: Hacky way of ignoring comments. Have them in a separate channel?
	// or just add them to a separate structure for future referencing:
	// comments.At(line, column) -> comment
	for {
		item := parser.lexer.NextItem()
		if item.Type != lexing.ItemComment {
			return item
		}
	}
}
