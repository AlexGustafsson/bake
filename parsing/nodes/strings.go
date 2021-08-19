package nodes

import (
	"fmt"
	"strings"
)

type InterpretedString struct {
	NodeType
	Range
	Content string
}

type RawString struct {
	NodeType
	Range
	Content string
}

func CreateInterpretedString(r Range, content string) *InterpretedString {
	return &InterpretedString{
		NodeType: NodeTypeInterpretedString,
		Range:    r,
		Content:  content,
	}
}

func CreateRawString(r Range, content string) *RawString {
	return &RawString{
		NodeType: NodeTypeRawString,
		Range:    r,
		Content:  content,
	}
}

func (node *RawString) String() string {
	return node.Content
}

func (node *InterpretedString) String() string {
	return node.Content
}

func (node *InterpretedString) DotString() string {
	return fmt.Sprintf("\"%p\" [label=\"interpreted string '%s'\"];\n", node, strings.ReplaceAll(node.Content, "\"", "\\\""))
}

func (node *RawString) DotString() string {
	escaped := strings.ReplaceAll(node.Content, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "\n", "\\n")
	return fmt.Sprintf("\"%p\" [label=\"raw string '%s'\"];\n", node, escaped)
}
