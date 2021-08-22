package semantics

// Scope is a scope within a program
type Scope struct {
	// ParentScope is the scope above this, if any (may be nil)
	ParentScope *Scope
	SymbolTable *SymbolTable
}

func CreateScope(parent *Scope) *Scope {
	return &Scope{
		ParentScope: parent,
		SymbolTable: CreateSymbolTable(),
	}
}

// LookupByName checks for a symbol by name up the scope tree
func (scope *Scope) LookupByName(name string) (*Symbol, bool) {
	symbol, ok := scope.SymbolTable.LookupByName(name)
	if !ok && scope.ParentScope != nil {
		symbol, ok = scope.ParentScope.LookupByName(name)
	}

	return symbol, ok
}
