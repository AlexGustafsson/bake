package runtime

import "fmt"

type DefaultRuntime struct {
}

func CreateDefaultRuntime() *DefaultRuntime {
	return &DefaultRuntime{}
}

func (runtime *DefaultRuntime) Add(left *Value, right *Value) *Value {
	if left.Type != right.Type {
		panic(fmt.Errorf("cannot add two values of different types"))
	}

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
	if left.Type != right.Type {
		panic(fmt.Errorf("cannot subtract two values of different types"))
	}

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
	if left.Type != right.Type {
		panic(fmt.Errorf("cannot multiply two values of different types"))
	}

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
	if left.Type != right.Type {
		panic(fmt.Errorf("cannot divide two values of different types"))
	}

	switch left.Type {
	case ValueTypeNumber:
		return &Value{
			Type:  ValueTypeNumber,
			Value: left.Value.(int) / right.Value.(int),
		}
	default:
		panic(fmt.Errorf("cannot divide values of type %s", left.Type))
	}
}

func (runtime *DefaultRuntime) DeclareVariable(identifier string, value *Value) {
	panic(fmt.Errorf("variable declarations are not implemented"))
}
