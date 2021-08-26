package ast

type ShellStatement struct {
	baseNode
	Multiline bool
	Parts     []Node
}

func CreateShellStatement(r *Range, multiline bool, parts []Node) *ShellStatement {
	return &ShellStatement{
		baseNode: baseNode{
			nodeType:  NodeTypeShellStatement,
			nodeRange: r,
		},
		Multiline: multiline,
		Parts:     parts,
	}
}
