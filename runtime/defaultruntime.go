package runtime

import (
	"fmt"
	"os/exec"
)

type DefaultRuntime struct {
	scope *Scope
}

func CreateDefaultRuntime() *DefaultRuntime {
	return &DefaultRuntime{
		scope: CreateScope(nil),
	}
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
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
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
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
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
		panic(fmt.Errorf("invalid operation for type %s", left.Type))
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

func (runtime *DefaultRuntime) Modulo(left *Value, right *Value) *Value {
	runtime.assertSameType(left, right)

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) % right.Value.(int),
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

func (runtime *DefaultRuntime) NotEquals(left *Value, right *Value) *Value {
	if left.Type != right.Type {
		return &Value{
			Type:  ValueTypeBool,
			Value: false,
		}
	}

	equal := false

	switch left.Type {
	case ValueTypeNumber:
		equal = left.Value.(int) != right.Value.(int)
	case ValueTypeString:
		equal = left.Value.(string) != right.Value.(string)
	case ValueTypeBool:
		equal = left.Value.(bool) != right.Value.(bool)
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

func (runtime *DefaultRuntime) Not(operand *Value) *Value {
	switch operand.Type {
	case ValueTypeBool:
		return &Value{
			Type:  ValueTypeBool,
			Value: !operand.Value.(bool),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", operand.Type))
	}
}

func (runtime *DefaultRuntime) Negative(operand *Value) *Value {
	switch operand.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: -operand.Value.(int),
		}
	default:
		panic(fmt.Errorf("invalid operation for type %s", operand.Type))
	}
}

func (runtime *DefaultRuntime) Shell(script string) {
	cmd := exec.Command("/bin/bash", "-c", script)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO: Use streams for realtime output
	fmt.Println(string(stdout))
}

func (runtime *DefaultRuntime) Define(identifier string, value *Value) {
	runtime.scope.Define(identifier, value)
}

func (runtime *DefaultRuntime) Resolve(identifier string) *Value {
	value := runtime.scope.Lookup(identifier)
	if value == nil {
		panic(fmt.Errorf("'%s' is undefined", identifier))
	}

	return value
}

func (runtime *DefaultRuntime) SetScope(scope *Scope) {
	runtime.scope = scope
}

func (runtime *DefaultRuntime) Scope() *Scope {
	return runtime.scope
}

func (runtime *DefaultRuntime) PushScope() {
	runtime.scope = runtime.scope.CreateScope()
}

func (runtime *DefaultRuntime) PopScope() {
	if runtime.scope.ParentScope == nil {
		panic(fmt.Errorf("cannot pop outside global scope"))
	}

	runtime.scope = runtime.scope.ParentScope
}

func (runtime *DefaultRuntime) assertSameType(left *Value, right *Value) {
	if left.Type != right.Type {
		panic(fmt.Errorf("invalid operation between two values of different types"))
	}
}
