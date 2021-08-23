package semantics

import "github.com/AlexGustafsson/bake/ast"

// Symbol is a semantic unit
type Symbol struct {
	// Name is the name of the symbol, such as the identifier of a declared variable
	Name string
	// Trait is the behavior of a symbol
	Trait Trait
	// ArgumentCount is the number of arguments the symbol may take if called
	ArgumentCount int
	// references is the set of lines on which the symbol is referenced
	references map[int]bool
	// Node is the AST node the symbol corresponds to
	Node ast.Node
}

func CreateSymbol(name string, trait Trait, node ast.Node) *Symbol {
	return &Symbol{
		Name:       name,
		Trait:      trait,
		Node:       node,
		references: make(map[int]bool),
	}
}

func (symbol *Symbol) AddReference(line int) {
	symbol.references[line] = true
}
