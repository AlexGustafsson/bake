package runtime

import (
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
)

type Program struct {
	Input    string
	Delegate Delegate
	Builtins map[string]*Builtin
}

// Builtin is a globally available value
type Builtin struct {
	Identifier string
	Symbol     *semantics.Symbol
	Value      *Value
}

func CreateProgram(input string, delegate Delegate) *Program {
	return &Program{
		Input:    input,
		Delegate: delegate,
		Builtins: make(map[string]*Builtin),
	}
}

// Run executes a program
func (program *Program) Run() []error {
	source, err := parsing.Parse(program.Input)
	if err != nil {
		return []error{err}
	}

	rootScope, errs := semantics.Build(source)
	if len(errs) > 0 {
		return errs
	}

	program.mountBuiltins(rootScope)

	errs = semantics.Validate(source, rootScope)
	if len(errs) > 0 {
		return errs
	}

	engine := CreateEngine(program.Delegate)
	err = engine.Evaluate(source)
	if err != nil {
		return []error{err}
	}

	return []error{}
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

// mountBuiltins declares symbols in the global scope and defines the values in the delegate's scope
func (program *Program) mountBuiltins(rootScope *semantics.Scope) {
	for _, builtin := range program.Builtins {
		rootScope.SymbolTable.Insert(builtin.Symbol)
		program.Delegate.Define(builtin.Identifier, builtin.Value)
	}
}
