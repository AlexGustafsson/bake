package builtins

import (
	"fmt"

	"github.com/AlexGustafsson/bake/runtime"
)

func registerRange(program *runtime.Program) {
	program.DefineBuiltinFunction("range", 2, func(engine *runtime.Engine, arguments []*runtime.Value) *runtime.Value {
		if arguments[0].Type != runtime.ValueTypeNumber || arguments[1].Type != runtime.ValueTypeNumber {
			panic(fmt.Errorf("bad arguments - expected numbers"))
		}

		start := arguments[0].Value.(int)
		stop := arguments[1].Value.(int)

		if stop < start {
			panic(fmt.Errorf("stop cannot be less than start"))
		}

		elements := make([]*runtime.Value, stop-start)
		for i := 0; i < stop-start; i++ {
			elements[i] = &runtime.Value{Type: runtime.ValueTypeNumber, Value: start + i}
		}

		return &runtime.Value{Type: runtime.ValueTypeArray, Value: elements}
	})
}
