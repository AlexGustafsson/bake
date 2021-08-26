package runtime

// Delegate handles runtime-specific evaluations. Functions may panic
type Delegate interface {
	Add(left *Value, b *Value) *Value
	Subtract(left *Value, b *Value) *Value
	Multiply(left *Value, b *Value) *Value
	Divide(left *Value, b *Value) *Value

	DeclareVariable(identifier string, value *Value)
}
