package semantics

// SymbolTable keeps track of symbols within a scope
type SymbolTable struct {
	symbolsByName map[string]*Symbol
}

func CreateSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbolsByName: make(map[string]*Symbol),
	}
}

// LookupByName checks for a symbol by name
func (table *SymbolTable) LookupByName(name string) (*Symbol, bool) {
	symbol, ok := table.symbolsByName[name]
	return symbol, ok
}

func (table *SymbolTable) Insert(symbol *Symbol) {
	table.symbolsByName[symbol.Name] = symbol
}
