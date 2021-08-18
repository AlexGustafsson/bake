package parsing

import (
	"container/list"
	"fmt"
	"runtime"
	"strings"

	"github.com/AlexGustafsson/bake/lexing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
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

func Parse(input string) (*nodes.SourceFile, error) {
	parser := CreateParser()
	parser.input = input
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

func (parser *Parser) tokenErrorf(item lexing.Item, format string, args ...interface{}) {
	// TODO: make less memory intensive
	lines := strings.Split(parser.input, "\n")
	line := lines[item.Line]

	var builder strings.Builder

	fmt.Fprintf(&builder, "\033[1mfile.bke:%d:%d\033[0m: \033[31;1merror\033[0m: ", item.Line, item.Column)
	fmt.Fprintf(&builder, format, args...)
	builder.WriteRune('\n')

	start := line[0:item.Column]
	end := line[item.Column+len(item.Value):]
	fmt.Fprintf(&builder, "%s\033[31;1m%s\033[0m%s\n", start, item.Value, end)
	builder.WriteString(strings.Repeat(" ", len(start)))
	fmt.Fprintf(&builder, "\033[31;1m^%s\033[0m\n", strings.Repeat("~", len(item.Value)-1))

	parser.errorf(builder.String())
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
