package lexing

import (
	"fmt"
	"strings"
)

type Item struct {
	Type    ItemType
	Start   int
	Value   string
	Message string
	Line    int
}

func (item Item) String() string {
	return item.Value
}

func (item Item) DebugString() string {
	// Escape newlines and tabs
	formattedValue := strings.ReplaceAll(strings.ReplaceAll(item.Value, "\n", "\\n"), "\t", "\\t")

	if item.Message != "" {
		return fmt.Sprintf("[item @%d:%d - %s \"%s\" - '%s']", item.Line, item.Start, item.Type.String(), item.Message, formattedValue)
	}

	return fmt.Sprintf("[item @%d:%d - %s '%s']", item.Line, item.Start, item.Type.String(), formattedValue)
}
