package parsing

import "github.com/AlexGustafsson/bake/parsing/nodes"

type SyntaxTree struct {
	Root nodes.Node
}

func CreateSyntaxTree(root nodes.Node) *SyntaxTree {
	return &SyntaxTree{
		Root: root,
	}
}
