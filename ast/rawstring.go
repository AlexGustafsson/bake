package ast

type RawString struct {
	baseNode
	Content string
}

func CreateRawString(r *Range, content string) *RawString {
	return &RawString{
		baseNode: baseNode{
			nodeType:  NodeTypeRawString,
			nodeRange: r,
		},
		Content: content,
	}
}

func (node *RawString) String() string {
	return node.Content
}
