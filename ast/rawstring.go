package ast

type RawString struct {
	NodeType
	Range
	Content string
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
