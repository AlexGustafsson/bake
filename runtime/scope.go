package runtime

type Scope struct {
	ParentScope *Scope
	Values      map[string]*Value
}

func CreateScope(parentScope *Scope) *Scope {
	return &Scope{
		ParentScope: parentScope,
		Values:      make(map[string]*Value),
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

func (scope *Scope) CreateScope() *Scope {
	return CreateScope(scope)
}
