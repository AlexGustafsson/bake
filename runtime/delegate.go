package runtime

// Delegate handles runtime-specific evaluations. Functions may panic
type Delegate interface {
	Add(left *Value, right *Value) *Value
	Subtract(left *Value, right *Value) *Value
	Multiply(left *Value, right *Value) *Value
	Divide(left *Value, right *Value) *Value
	Modulo(left *Value, right *Value) *Value

	Equals(left *Value, right *Value) *Value
	NotEquals(left *Value, right *Value) *Value
	GreaterThan(left *Value, right *Value) *Value
	GreaterThanOrEqual(left *Value, right *Value) *Value
	LessThan(left *Value, right *Value) *Value
	LessThanOrEqual(left *Value, right *Value) *Value

	And(left *Value, right *Value) *Value
	Or(left *Value, right *Value) *Value

	Not(operand *Value) *Value
	Negative(operand *Value) *Value

	Shell(script string)
	ShellFormat(value *Value) string

	ResolveEscapes(value string) string

	Define(identifier string, value *Value)
	Resolve(identifier string) *Value
	SetScope(scope *Scope)
	Scope() *Scope
	PushScope(scopeType ScopeType)
	PopScope()
}
