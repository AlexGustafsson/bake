package runtime

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type DefaultRuntime struct {
	scope *Scope
}

func CreateDefaultRuntime() *DefaultRuntime {
	return &DefaultRuntime{
		scope: CreateScope(nil, ScopeTypeGlobal),
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
	cmd := exec.Command("/bin/bash")
	var stdout, stderr bytes.Buffer
	// Always end the script with a newline to ensure it's executed
	if !strings.HasSuffix(script, "\n") {
		script += "\n"
	}
	cmd.Stdin = strings.NewReader(script)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	context := runtime.Resolve("context").Value.(Object)
	if shellValue, ok := context["shell"]; ok {
		shell := shellValue.Value.(Object)
		shell["stdout"].Value = strings.TrimSpace(stdout.String())
		shell["stderr"].Value = strings.TrimSpace(stderr.String())
		shell["status"].Value = cmd.ProcessState.ExitCode()
	}
}

func (runtime *DefaultRuntime) ResolveEscapes(value string) string {
	value = strings.ReplaceAll(value, "\\\"", "\"")
	value = strings.ReplaceAll(value, "\\t", "\t")
	value = strings.ReplaceAll(value, "\\n", "\n")
	value = strings.ReplaceAll(value, "\\$", "$")
	return value
}

func (runtime *DefaultRuntime) ShellFormat(value *Value) string {
	switch value.Type {
	case ValueTypeNone:
		return ""
	case ValueTypeNumber:
		return fmt.Sprintf("%d", value.Value)
	case ValueTypeString:
		return fmt.Sprintf("%s", value.Value)
	case ValueTypeBool:
		return fmt.Sprintf("%t", value.Value)
	case ValueTypeArray:
		array := value.Value.(Array)
		var builder strings.Builder
		for i, x := range array {
			if i > 0 {
				builder.WriteString(" ")
			}
			builder.WriteString(runtime.ShellFormat(x))
		}
		return builder.String()
	default:
		panic(fmt.Errorf("cannot format value '%s' for current shell", value.String()))
	}
}

func (runtime *DefaultRuntime) Index(left *Value, right *Value) *Value {
	if left.Type == ValueTypeObject {
		if right.Type == ValueTypeString {
			object := left.Value.(Object)
			if value, ok := object[right.Value.(string)]; ok {
				return value
			} else {
				return &Value{Type: ValueTypeNone}
			}
		} else {
			panic(fmt.Errorf("an object can only be indexed with a string"))
		}
	} else if left.Type == ValueTypeArray {
		if right.Type == ValueTypeNumber {
			index := right.Value.(int)
			array := left.Value.(Array)
			if index >= 0 && index < len(array) {
				return array[index]
			} else {
				panic(fmt.Errorf("index is out of bounds"))
			}
		} else {
			panic(fmt.Errorf("an array can only be indexed with a number"))
		}
	} else {
		panic(fmt.Errorf("cannot index a value of type '%s'", left.Type))
	}
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

func (runtime *DefaultRuntime) PushScope(scopeType ScopeType) {
	runtime.scope = runtime.scope.CreateScope(scopeType)
	if scopeType == ScopeTypeFunction || scopeType == ScopeTypeRuleFunction || scopeType == ScopeTypeRule {
		context := make(Object)

		shell := make(Object)
		shell["stdout"] = &Value{
			Type:  ValueTypeString,
			Value: "",
		}
		shell["stderr"] = &Value{
			Type:  ValueTypeString,
			Value: "",
		}
		shell["status"] = &Value{
			Type:  ValueTypeNumber,
			Value: 0,
		}

		context["shell"] = &Value{
			Type:  ValueTypeObject,
			Value: shell,
		}

		if scopeType == ScopeTypeRule || scopeType == ScopeTypeRuleFunction {
			context["input"] = &Value{
				Type:  ValueTypeArray,
				Value: make(Array, 0),
			}

			context["output"] = &Value{
				Type:  ValueTypeArray,
				Value: make(Array, 0),
			}
		}

		value := &Value{Type: ValueTypeObject, Value: context}
		runtime.Define("context", value)
	}
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
