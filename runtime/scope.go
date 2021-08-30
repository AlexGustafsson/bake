package runtime

type ScopeType int

const (
	ScopeTypeFunction ScopeType = iota
	ScopeTypeRuleFunction
	ScopeTypeRule
	ScopeTypeGeneric
	ScopeTypeGlobal
)

type Scope struct {
	ParentScope *Scope
	Values      map[string]*Value
	Type        ScopeType
}

func CreateScope(parentScope *Scope, scopeType ScopeType) *Scope {
	return &Scope{
		ParentScope: parentScope,
		Values:      make(map[string]*Value),
		Type:        scopeType,
	}
}

// Lookup looks up an identifier upward the scope tree. Returns nil if not found
func (scope *Scope) Lookup(identifier string) *Value {
	if value, ok := scope.Values[identifier]; ok {
		return value
	}

	if scope.ParentScope != nil {
		return scope.ParentScope.Lookup(identifier)
	}

	return nil
}

func (scope *Scope) Define(identifier string, value *Value) {
	scope.Values[identifier] = value
}

func (scope *Scope) CreateScope(scopeType ScopeType) *Scope {
	return CreateScope(scope, scopeType)
}
