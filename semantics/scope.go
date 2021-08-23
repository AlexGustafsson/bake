package semantics

// Scope is a scope within a program
type Scope struct {
	// ParentScope is the scope above this, if any (may be nil)
	ParentScope *Scope
	// ChildScopes are the scopes contained within this, in the order they're declared
	ChildScopes []*Scope
	// SymbolTable is the table of symbols defined within the scope
	SymbolTable *SymbolTable
}

func CreateScope(parent *Scope) *Scope {
	return &Scope{
		ParentScope: parent,
		ChildScopes: make([]*Scope, 0),
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

func (scope *Scope) CreateScope() *Scope {
	child := CreateScope(scope)
	scope.ChildScopes = append(scope.ChildScopes, child)
	return child
}
