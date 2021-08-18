package lexing

import (
	"fmt"
	"strings"
)

type Item struct {
	// Type is the type of the lexical item
	Type ItemType
	// Value is the string value of the token
	Value string
	// Message is an optional string for the token
	Message string
	// Start is the offset to the start of the token in bytes
	Start int
	// Line is the zero-based line of the input on which the token is found
	Line int
	// Column is the zero-based column (in runes) of the line in which the token is found
	Column int
}

func (item Item) String() string {
	return item.Value
}

func (item Item) DebugString() string {
	// Escape newlines and tabs
	formattedValue := strings.ReplaceAll(strings.ReplaceAll(item.Value, "\n", "\\n"), "\t", "\\t")

	if item.Message != "" {
		return fmt.Sprintf("[item @%d:%d - %s \"%s\" - '%s']", item.Line+1, item.Column+1, item.Type.String(), item.Message, formattedValue)
	}

	return fmt.Sprintf("[item @%d:%d - %s '%s']", item.Line+1, item.Column+1, item.Type.String(), formattedValue)
}
