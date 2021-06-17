package lexing

import "fmt"

type Item struct {
	Type    ItemType
	Start   int
	Value   string
	Message string
	Line    int
}

func (item Item) String() string {
	if item.Message != "" {
		return fmt.Sprintf("[item @%d:%d - %d \"%s\" - '%s']", item.Line, item.Start, item.Type, item.Message, item.Value)
	}

	return fmt.Sprintf("[item @%d:%d - %d '%s']", item.Line, item.Start, item.Type, item.Value)
}
