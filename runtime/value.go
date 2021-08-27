package runtime

import (
	"fmt"

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
)

type Value struct {
	Type  ValueType
	Value interface{}
}

func (value *Value) String() string {
	if value.Type == ValueTypeNone {
		return "none"
	}

	switch cast := value.Value.(type) {
	case int:
		return fmt.Sprintf("%d", cast)
	case string:
		return cast
	case bool:
		return fmt.Sprintf("%t", cast)
	}

	return "?"
}

type Rule struct {
	Outputs      []string
	Dependencies []*Value
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
