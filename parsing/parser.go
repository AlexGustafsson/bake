package parsing

import (
	"container/list"
	"fmt"
	"runtime"

	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
)

type Parser struct {
	lexer  *lexing.Lexer
	peeked *list.List
}

func CreateParser() *Parser {
	return &Parser{
		lexer:  nil,
		peeked: list.New(),
	}
}

func Parse(input string) (*nodes.SourceFile, error) {
	parser := CreateParser()
	return parser.Parse(input)
}

func (parser *Parser) Parse(input string) (_ *nodes.SourceFile, err error) {
	parser.lexer = lexing.Lex(input)

	// Recover from panics caused when parsing
	defer parser.recover(&err)

	sourceFile, err := parseSourceFile(parser)
	if err != nil {
		return nil, err
	}

	return sourceFile, nil
}

func (parser *Parser) errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
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
		parser.errorf("line %d: expected %s, got %s", item.Line, itemType.String(), item.Type.String())
	}
	return item
}

func (parser *Parser) expect(itemType lexing.ItemType) (_ lexing.Item, ok bool) {
	item := parser.nextItem()

	if item.Type == lexing.ItemError {
		parser.errorf("line %d: %s", item.Line, item.Message)
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

	return parser.lexer.NextItem()
}
