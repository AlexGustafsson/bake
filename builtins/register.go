package builtins

import "github.com/AlexGustafsson/bake/runtime"

func Register(program *runtime.Program) {
	registerPrint(program)
	registerRange(program)
	registerEnv(program)
}
