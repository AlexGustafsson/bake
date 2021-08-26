package runtime

// Delegate handles runtime-specific evaluations. Functions may panic
type Delegate interface {
	Add(left *Value, right *Value) *Value
	Subtract(left *Value, right *Value) *Value
	Multiply(left *Value, right *Value) *Value
	Divide(left *Value, right *Value) *Value

	Equals(left *Value, right *Value) *Value
	GreaterThan(left *Value, right *Value) *Value
	GreaterThanOrEqual(left *Value, right *Value) *Value
	LessThan(left *Value, right *Value) *Value
	LessThanOrEqual(left *Value, right *Value) *Value

	DeclareVariable(identifier string, value *Value)
}
