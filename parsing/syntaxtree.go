package parsing

type SyntaxTree struct {
	Imports []string
}

func CreateSyntaxTree() *SyntaxTree {
	return &SyntaxTree{
		Imports: make([]string, 0),
	}
}
