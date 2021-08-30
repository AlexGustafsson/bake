package semantics

//go:generate stringer -type=Trait
// Trait describes the behavior of a symbol
type Trait uint64

// TraitNone indiciates an empty trait
const TraitNone Trait = 0

// definedTraits are the number of available traits (excluding TraitNone)
const definedTraits = 3

const (
	// TraitNumeric indicates that the symbol may be used for numeric operations
	TraitNumeric Trait = 1 << iota
	// TraitCallable indicates that the symbol may be called
	TraitCallable
	// TraitAny indicates that the symbol behaves like any trait (likely determined at runtime for non-typed variables)
	TraitAny
	// TraitAlias indicates that the symbol is an alias for other rules, aliases or functions
	TraitAlias
	// TraitImport indicates that the symbol is the root of an import package
	TraitImport
	// TraitString indicates that the symbol may be treated as a string
	TraitString
	// TraitObject indiciates that the symbol may be treated as a string
	TraitObject
)

// Has checks whether or not a trait has the specified trait
func (trait Trait) Has(other Trait) bool {
	return trait&other != 0
}

func (trait Trait) Strings() []string {
	strings := make([]string, 0)
	// Strings returns the strings of all contained traits
	for i := 0; i < definedTraits; i++ {
		var other Trait = 1 << i
		if trait.Has(other) {
			strings = append(strings, other.String())
		}
	}
	return strings
}
