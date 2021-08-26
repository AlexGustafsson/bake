package ast

type SourceFile struct {
	baseNode
	Nodes []Node
}

func CreateSourceFile(r *Range) *SourceFile {
	return &SourceFile{
		baseNode: baseNode{
			nodeType:  NodeTypeSourceFile,
			nodeRange: r,
		},
		Nodes: make([]Node, 0),
	}
}
