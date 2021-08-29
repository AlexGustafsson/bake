package runtime

import (
	"github.com/AlexGustafsson/bake/semantics"
)

// Program is an entire Bake program
type Program struct {
	// Main is the main entrypoint file of the program, which will be used as the global scope
	MainModule *Module
	// Builtins are globally available values
	Builtins map[string]*Builtin
	// ImportedModules are all the imported modules by their package
	ImportedModules    map[string]*Module
	DependencyResolver *DependencyResolver
}

// Builtin is a globally available value
type Builtin struct {
	Identifier string
	Symbol     *semantics.Symbol
	Value      *Value
}

func CreateProgram(input string) *Program {
	return &Program{
		MainModule: CreateModule(input),
		Builtins:   make(map[string]*Builtin),
	}
}

// Parse parses the source of the program and its dependencies
func (program *Program) Parse() error {
	return program.MainModule.Parse()
}

// ResolveImports fetches all imports to the root director
func (program *Program) ResolveImports(root string) []error {
	program.DependencyResolver = CreateDependencyResolver(root)
	return program.DependencyResolver.Resolve(program.MainModule.Imports())
}

// BuildSymbols builds the root semantic scope of the parsed program
func (program *Program) BuildSymbols() []error {
	errs := program.MainModule.BuildSymbols()
	for _, resolvedPackage := range program.DependencyResolver.ResolvedPackages {
		for _, module := range resolvedPackage.Modules {
			errs = append(errs, module.BuildSymbols()...)
		}
	}
	return errs
}

// DefineBuiltinSymbols declares symbols in the global scope
func (program *Program) DefineBuiltinSymbols() {
	for _, builtin := range program.Builtins {
		program.MainModule.RootScope.SymbolTable.Insert(builtin.Symbol)
	}
}

// DefineBuiltinSymbols defines the values in the delegate's scope
func (program *Program) DefineBuiltinValues(delegate Delegate) {
	for _, builtin := range program.Builtins {
		delegate.Define(builtin.Identifier, builtin.Value)
	}
}

// Validate validates the semantics of the program
func (module *Program) Validate() []error {
	return module.MainModule.Validate()
	// TODO: Ensure that this validation is done for all values that are used by the program, dependency or not
}

// DefineBuiltinFunction defines as new, globally available function
func (program *Program) DefineBuiltinFunction(identifier string, arguments int, handler FunctionHandler) {
	symbol := semantics.CreateSymbol(identifier, semantics.TraitCallable, nil)
	symbol.ArgumentCount = arguments

	function := &Function{
		Handler: handler,
	}

	value := &Value{
		Type:  ValueTypeFunction,
		Value: function,
	}

	program.Builtins[identifier] = &Builtin{identifier, symbol, value}
}

// Run executes the program using the specified delegate
func (program *Program) Run(delegate Delegate) error {
	engine := CreateEngine(delegate)
	return program.MainModule.Evaluate(engine, delegate)
}
