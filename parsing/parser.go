package parsing

import (
	"fmt"
	"runtime"

	"github.com/AlexGustafsson/bake/lexing"
)

type Parser struct {
	lexer      *lexing.Lexer
	syntaxTree *SyntaxTree
}

func CreateParser() *Parser {
	return &Parser{
		lexer:      nil,
		syntaxTree: &SyntaxTree{},
	}
}

func Parse(input string) (*SyntaxTree, error) {
	parser := CreateParser()
	return parser.Parse(input)
}

func (parser *Parser) Parse(input string) (_ *SyntaxTree, err error) {
	parser.lexer = lexing.Lex(input)

	// Recover from panics caused when parsing
	defer parser.recover(&err)

	root, err := parseRoot(parser)
	if err != nil {
		return nil, err
	}

	tree := CreateSyntaxTree(root)

	return tree, nil
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
