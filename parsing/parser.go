package parsing

import (
	"fmt"
	"sync"

	"github.com/AlexGustafsson/bake/lexing"
)

type Parser struct {
	lexer *lexing.Lexer
	wg    sync.WaitGroup
}

func Parse(input string) *Parser {
	parser := &Parser{
		lexer: lexing.Lex(input),
	}
	parser.wg.Add(1)
	go parser.run()
	return parser
}

func (parser *Parser) run() {
	for {
		item := parser.lexer.NextItem()
		if item.Type == lexing.ItemError {
			fmt.Printf("Got error token: %s\n", item)
			break
		} else if item.Type == lexing.ItemEndOfFile {
			fmt.Printf("Got EOF\n")
			break
		}

		fmt.Printf("Got token: %s\n", item)
	}
	parser.wg.Done()
}

func (parser *Parser) Wait() {
	parser.wg.Wait()
}
