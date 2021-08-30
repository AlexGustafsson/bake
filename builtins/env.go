package builtins

import (
	"os"
	"strings"

	"github.com/AlexGustafsson/bake/runtime"
	"github.com/AlexGustafsson/bake/semantics"
)

func registerEnv(program *runtime.Program) {
	// env is the globally available object containing environment variables
	env := make(runtime.Object)

	for _, variable := range os.Environ() {
		pair := strings.SplitN(variable, "=", 2)
		value := &runtime.Value{Type: runtime.ValueTypeString, Value: pair[1]}
		env[pair[0]] = value
	}

	program.DefineBuiltinValue("env", runtime.ValueTypeObject, semantics.TraitObject, env)
}
