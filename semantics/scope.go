package semantics

import "github.com/AlexGustafsson/bake/ast"

// Scope is a scope within a program
type Scope struct {
	// ParentScope is the scope above this, if any (may be nil)
	ParentScope *Scope
	// ChildScopes are the scopes contained within this, in the order they're declared
	ChildScopes map[ast.Node]*Scope
	// SymbolTable is the table of symbols defined within the scope
	SymbolTable *SymbolTable
}

func CreateScope(parent *Scope) *Scope {
	return &Scope{
		ParentScope: parent,
		ChildScopes: make(map[ast.Node]*Scope),
		SymbolTable: CreateSymbolTable(),
	}
}

// LookupByName checks for a symbol by name up the scope tree
func (scope *Scope) LookupByName(name string) (*Symbol, *Scope, bool) {
	symbol, ok := scope.SymbolTable.LookupByName(name)
	foundScope := scope
	if !ok && scope.ParentScope != nil {
		symbol, foundScope, ok = scope.ParentScope.LookupByName(name)
	}

	return symbol, foundScope, ok
}

// LookupNode checks for a symbol by name up the scope tree
func (scope *Scope) LookupNode(node ast.Node) (*Symbol, *Scope, bool) {
	symbol, ok := scope.SymbolTable.LookupNode(node)
	foundScope := scope
	if !ok && scope.ParentScope != nil {
		symbol, foundScope, ok = scope.ParentScope.LookupNode(node)
	}

	return symbol, foundScope, ok
}

func (scope *Scope) CreateScope(node ast.Node) *Scope {
	child := CreateScope(scope)
	scope.ChildScopes[node] = child
	return child
}
