package ast

import (
	"strings"
)

type EvaluatedString struct {
	NodeType
	Range
	Parts []Node
}

type StringPart struct {
	NodeType
	Range
	Content string
}

func CreateEvaluatedString(r Range, parts []Node) *EvaluatedString {
	return &EvaluatedString{
		NodeType: NodeTypeEvaluatedString,
		Range:    r,
		Parts:    parts,
	}
}

func CreateStringPart(r Range, content string) *StringPart {
	return &StringPart{
		NodeType: NodeTypeStringPart,
		Range:    r,
		Content:  content,
	}
}

func (node *EvaluatedString) String() string {
	var builder strings.Builder

	builder.WriteRune('"')

	for _, part := range node.Parts {
		builder.WriteString(part.String())
	}

	builder.WriteRune('"')

	return builder.String()
}

func (node *StringPart) String() string {
	return node.Content
}
