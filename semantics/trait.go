package semantics

// Trait describes the behavior of a symbol
type Trait uint64

// TraitNone indiciates an empty trait
const TraitNone Trait = 0

const (
	// TraitNumeric indicates that the symbol may be used for numeric operations
	TraitNumeric Trait = 1 << iota
)

// Has checks whether or not a trait has the specified trait
func (trait Trait) Has(other Trait) bool {
	return trait&other != 0
}
