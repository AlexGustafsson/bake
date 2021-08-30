package runtime

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/ast"
)

//go:generate stringer -type=ValueType
type ValueType int

const (
	ValueTypeNumber ValueType = iota
	ValueTypeString
	ValueTypeBool
	ValueTypeFunction
	ValueTypeRuleFunction
	ValueTypeRule
	ValueTypeNone
	ValueTypeArray
	ValueTypeAlias
	ValueTypeObject
)

type Value struct {
	Type     ValueType
	Value    interface{}
	Exported bool
}

func (value *Value) String() string {
	if value.Type == ValueTypeNone {
		return "none"
	}

	switch cast := value.Value.(type) {
	case int:
		return fmt.Sprintf("%d", cast)
	case string:
		return fmt.Sprintf("\"%s\"", cast)
	case bool:
		return fmt.Sprintf("%t", cast)
	case Array:
		var builder strings.Builder
		builder.WriteRune('[')
		for i, x := range cast {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(x.String())
		}
		builder.WriteRune(']')
		return builder.String()
	case Object:
		var builder strings.Builder
		builder.WriteRune('{')
		i := 0
		for key, value := range cast {
			if i > 0 {
				builder.WriteString(", ")
			}
			fmt.Fprintf(&builder, "'%s': %s", key, value.String())
			i++
		}
	}

	return fmt.Sprintf("<unknown %s>", value.Type)
}

type Rule struct {
	Outputs      []string
	Dependencies []*Value
	// Block is the AST block that will be executed on invocation
	Block *ast.Block
}

type FunctionHandler func(engine *Engine, arguments []*Value) *Value

type Function struct {
	// Arguments to define before executing the block. Only used when evaluating the Block
	Arguments []string
	// Block is the AST block that will be executed on invocation. Block and Handler are mutually exclusive.
	Block *ast.Block
	// Handler is the callback that should handle the invocation. Block and Handler are mutually exclusive.
	Handler        FunctionHandler
	IsRuleFunction bool
}

type Alias struct {
	Dependencies []*Value
}

type Object map[string]*Value
type Array []*Value
