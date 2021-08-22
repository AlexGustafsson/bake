package semantics

// Symbol is a semantic unit
type Symbol struct {
	Name  string
	Trait Trait
}

func CreateSymbol(name string, trait Trait) *Symbol {
	return &Symbol{
		Name: name,
	}
}
