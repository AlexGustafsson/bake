package ast

type EvaluatedString struct {
	baseNode
	Parts []Node
}

type StringPart struct {
	baseNode
	Content string
}

func CreateEvaluatedString(r *Range, parts []Node) *EvaluatedString {
	return &EvaluatedString{
		baseNode: baseNode{
			nodeType:  NodeTypeEvaluatedString,
			nodeRange: r,
		},
		Parts: parts,
	}
}

func CreateStringPart(r *Range, content string) *StringPart {
	return &StringPart{
		baseNode: baseNode{
			nodeType:  NodeTypeStringPart,
			nodeRange: r,
		},
		Content: content,
	}
}
