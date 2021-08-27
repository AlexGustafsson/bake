package builtins

import (
	"fmt"

	"github.com/AlexGustafsson/bake/runtime"
)

func registerPrint(program *runtime.Program) {
	program.DefineBuiltinFunction("print", -1, func(engine *runtime.Engine, arguments []*runtime.Value) *runtime.Value {
		for i, argument := range arguments {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(argument)
		}
		fmt.Println()
		return nil
	})
}
