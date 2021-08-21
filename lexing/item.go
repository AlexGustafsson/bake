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
	Range   Range
}

func (item Item) String() string {
	return item.Value
}

func (item Item) DebugString() string {
	// Escape newlines and tabs
	formattedValue := strings.ReplaceAll(strings.ReplaceAll(item.Value, "\n", "\\n"), "\t", "\\t")

	if item.Message != "" {
		return fmt.Sprintf("[%s %s - \"%s\" - '%s']", item.Type, item.Range, item.Message, formattedValue)
	}

	return fmt.Sprintf("[%s %s - '%s']", item.Type, item.Range, formattedValue)
}

func (item Item) IsKeyword() bool {
	switch item.Type {
	case ItemKeywordPackage, ItemKeywordImport, ItemKeywordFunc, ItemKeywordRule, ItemKeywordExport, ItemKeywordIf, ItemKeywordElse, ItemKeywordReturn, ItemKeywordLet, ItemKeywordShell, ItemKeywordAlias:
		return true
	default:
		return false
	}
}
