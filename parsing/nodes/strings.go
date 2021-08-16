package nodes

type InterpretedString struct {
	NodeType
	NodePosition
	Content string
}

type RawString struct {
	NodeType
	NodePosition
	Content string
}

func CreateInterpretedString(position NodePosition, content string) *InterpretedString {
	return &InterpretedString{
		NodeType:     NodeTypeInterpretedString,
		NodePosition: position,
		Content:      content,
	}
}

func CreateRawString(position NodePosition, content string) *RawString {
	return &RawString{
		NodeType:     NodeTypeRawString,
		NodePosition: position,
		Content:      content,
	}
}

func (node *RawString) String() string {
	return node.Content
}

func (node *InterpretedString) String() string {
	return node.Content
}
