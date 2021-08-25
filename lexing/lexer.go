package lexing

import (
	"fmt"
	"unicode/utf8"
)

type stateModifier func(lexer *Lexer) stateModifier

type LexMode int

const (
	ModeRoot LexMode = iota
	ModeEvaluatedString
	ModeShellString
	ModeMultilineShellString
)

type Lexer struct {
	input    string
	position int
	// startCharacter is the offset in the current line to the first rune of the current token
	startCharacter int
	// character is the offset in the current line to the current rune
	character int
	startLine int
	// lineLengths are the lengths of past lines in runes
	lineLengths       []int
	line              int
	start             int
	runeWidth         int
	items             chan Item
	parenthesesDepth  int
	substitutionDepth int

	Mode LexMode
}

func Lex(input string) *Lexer {
	lexer := &Lexer{
		input:       input,
		lineLengths: make([]int, 0),
		items:       make(chan Item),
		Mode:        ModeRoot,
	}
	go lexer.run()
	return lexer
}

func (lexer *Lexer) run() {
	// Push initial token
	lexer.Emit(ItemStartOfInput)

	// Run through the states
	for state := lexRoot; state != nil; {
		state = state(lexer)
	}

	close(lexer.items)
}

func (lexer *Lexer) Next() rune {
	if lexer.position >= len(lexer.input) {
		lexer.runeWidth = 0
		return eof
	}

	rune, runeWidth := utf8.DecodeRuneInString(lexer.input[lexer.position:])
	lexer.runeWidth = runeWidth
	lexer.position += runeWidth

	if rune == '\n' {
		lexer.line++
		lexer.lineLengths = append(lexer.lineLengths, lexer.character)
		lexer.character = 0
	} else {
		lexer.character++
	}

	return rune
}

func (lexer *Lexer) NextString(text string) bool {
	accepted := 0
	for _, expected := range text {
		actual := lexer.Next()
		if actual != expected {
			break
		}

		accepted++
	}

	ok := accepted == len(text)
	if !ok {
		lexer.BacktrackCount(accepted)
	}

	return ok
}

func (lexer *Lexer) Peek() rune {
	rune := lexer.Next()
	lexer.Backtrack()
	return rune
}

func (lexer *Lexer) Backtrack() {
	lexer.position -= lexer.runeWidth
	lexer.character--
	if lexer.runeWidth == 1 && lexer.input[lexer.position] == '\n' {
		lexer.line--
		stackSize := len(lexer.lineLengths)
		lexer.character = lexer.lineLengths[stackSize-1]
		lexer.lineLengths = lexer.lineLengths[:stackSize]
	}
}

func (lexer *Lexer) BacktrackCount(count int) {
	for i := 0; i < count; i++ {
		lexer.Backtrack()
	}
}

func (lexer *Lexer) Emit(itemType ItemType) {
	lexer.items <- Item{
		Type:  itemType,
		Value: lexer.input[lexer.start:lexer.position],
		Range: Range{
			Start: Position{
				Offset:    lexer.start,
				Line:      lexer.startLine,
				Character: lexer.startCharacter,
			},
			End: Position{
				Offset:    lexer.position,
				Line:      lexer.line,
				Character: lexer.character,
			},
		},
	}
	lexer.start = lexer.position
	lexer.startLine = lexer.line
	lexer.startCharacter = lexer.character
}

func (lexer *Lexer) EmitWithMessage(itemType ItemType, message string) {
	lexer.items <- Item{
		Type:    itemType,
		Value:   lexer.input[lexer.start:lexer.position],
		Message: message,
		Range: Range{
			Start: Position{
				Offset:    lexer.start,
				Line:      lexer.startLine,
				Character: lexer.startCharacter,
			},
			End: Position{
				Offset:    lexer.position,
				Line:      lexer.line,
				Character: lexer.character,
			},
		},
	}
	lexer.start = lexer.position
	lexer.startLine = lexer.line
	lexer.startCharacter = lexer.character
}

func (lexer *Lexer) NextItem() Item {
	return <-lexer.items
}

func (lexer *Lexer) NextNonWhitespaceItem(includeNewline bool) Item {
	for {
		item := <-lexer.items
		if item.Type != ItemWhitespace && !(includeNewline && item.Type == ItemNewline) {
			return item
		}
	}
}

func (lexer *Lexer) AllItems() []Item {
	items := make([]Item, 0)
	for item := range lexer.items {
		items = append(items, item)
	}
	return items
}

func (lexer *Lexer) Ignore() {
	lexer.start = lexer.position
	lexer.startLine = lexer.line
	lexer.startCharacter = lexer.character
}

func (lexer *Lexer) errorf(format string, args ...interface{}) {
	lexer.EmitWithMessage(ItemError, fmt.Sprintf(format, args...))
}
