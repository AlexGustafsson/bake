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
)

type Value struct {
	Type  ValueType
	Value interface{}
}

func (value *Value) String() string {
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

type Function struct {
	Arguments      []string
	Block          *ast.Block
	IsRuleFunction bool
}
