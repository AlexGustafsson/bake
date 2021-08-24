package semantics

import "github.com/AlexGustafsson/bake/ast"

// SymbolTable keeps track of symbols within a scope
type SymbolTable struct {
	symbolsByName           map[string]*Symbol
	symbolsByNode           map[ast.Node]*Symbol
	symbolsByInsertionOrder map[int]*Symbol
	symbolCount             int
}

func CreateSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbolsByName:           make(map[string]*Symbol),
		symbolsByNode:           make(map[ast.Node]*Symbol),
		symbolsByInsertionOrder: make(map[int]*Symbol),
	}
}

// LookupByName checks for a symbol by name
func (table *SymbolTable) LookupByName(name string) (*Symbol, bool) {
	symbol, ok := table.symbolsByName[name]
	return symbol, ok
}

// LookupNode checks for a symbol for the node
func (table *SymbolTable) LookupNode(node ast.Node) (*Symbol, bool) {
	symbol, ok := table.symbolsByNode[node]
	return symbol, ok
}

// Insert directly inserts a symbol
func (table *SymbolTable) Insert(symbol *Symbol) {
	if symbol.Name != "" {
		table.symbolsByName[symbol.Name] = symbol
	}

	if symbol.Node != nil {
		table.symbolsByNode[symbol.Node] = symbol
	}

	table.symbolsByInsertionOrder[table.symbolCount] = symbol
	table.symbolCount++
}

// Symbols returns all of the symbols in the order they were inserted
func (table *SymbolTable) Symbols() []*Symbol {
	symbols := make([]*Symbol, 0)
	for i := 0; i < table.symbolCount; i++ {
		symbols = append(symbols, table.symbolsByInsertionOrder[i])
	}
	return symbols
}
