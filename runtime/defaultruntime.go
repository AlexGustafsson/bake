package runtime

import "fmt"

type DefaultRuntime struct {
}

func CreateDefaultRuntime() *DefaultRuntime {
	return &DefaultRuntime{}
}

func (runtime *DefaultRuntime) Add(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) + right.Value.(int),
		}
	case ValueTypeString:
		return &Value{
			Type:  ValueTypeString,
			Value: left.Value.(string) + right.Value.(string),
		}
	default:
		panic(fmt.Errorf("cannot add values of type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) Subtract(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) - right.Value.(int),
		}
	default:
		panic(fmt.Errorf("cannot subtract values of type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) Multiply(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) * right.Value.(int),
		}
	default:
		panic(fmt.Errorf("cannot multiply values of type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) Divide(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) / right.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) Equals(left *Value, right *Value) *Value {
	if left.Type != right.Type {
		return &Value{
			Type:  ValueTypeBool,
			Value: false,
		}
	}

	equal := false

	switch left.Type {
	case ValueTypeNumber:
		equal = left.Value.(int) == right.Value.(int)
	case ValueTypeString:
		equal = left.Value.(string) == right.Value.(string)
	case ValueTypeBool:
		equal = left.Value.(bool) == right.Value.(bool)
	}

	return &Value{
		Type:  ValueTypeBool,
		Value: equal,
	}
}

func (runtime *DefaultRuntime) GreaterThan(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(int) > right.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) GreaterThanOrEqual(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(int) >= right.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) LessThan(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(int) < right.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) LessThanOrEqual(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(int) <= right.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) And(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeBool:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(bool) && right.Value.(bool),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) Or(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeBool:
		return &Value{
			Type:  ValueTypeBool,
			Value: left.Value.(bool) || right.Value.(bool),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) DeclareVariable(identifier string, value *Value) {
	panic(fmt.Errorf("variable declarations are not implemented"))
}

func (runtime *DefaultRuntime) assertSameType(left *Value, right *Value) {
	if left.Type != right.Type {
		panic(fmt.Errorf("invalid operation between two values of different types"))
	}
}
