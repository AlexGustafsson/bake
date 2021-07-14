package parsing

import (
	"fmt"
	"sync"

	"github.com/AlexGustafsson/bake/lexing"
)

type stateModifier func(parser *Parser) stateModifier

type Parser struct {
	lexer      *lexing.Lexer
	wg         sync.WaitGroup
	syntaxTree *SyntaxTree
}

func Parse(input string) *Parser {
	parser := &Parser{
		lexer:      lexing.Lex(input),
		syntaxTree: CreateSyntaxTree(),
	}
	parser.wg.Add(1)
	go parser.run()
	return parser
}

func (parser *Parser) run() {
	for state := parseRoot; state != nil; {
		state = state(parser)
	}

	parser.wg.Done()
}

func (parser *Parser) Wait() *SyntaxTree {
	parser.wg.Wait()
	return parser.syntaxTree
}

func (parser *Parser) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
