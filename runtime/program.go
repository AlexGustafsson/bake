package runtime

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
)

type Program struct {
	Input     string
	Source    *ast.Block
	RootScope *semantics.Scope
	Builtins  map[string]*Builtin
	Engine    *Engine
	Delegate  Delegate
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
		Builtins: make(map[string]*Builtin),
		Delegate: delegate,
		Engine:   CreateEngine(delegate),
	}
}

func (program *Program) Parse() error {
	source, err := parsing.Parse(program.Input)
	if err != nil {
		return err
	}

	program.Source = source
	return nil
}

func (program *Program) BuildSymbols() []error {
	rootScope, errs := semantics.Build(program.Source)
	if len(errs) > 0 {
		return errs
	}

	program.RootScope = rootScope
	return []error{}
}

// DefineBuiltinSymbols declares symbols in the global scope
func (program *Program) DefineBuiltinSymbols() {
	// TODO: Define before others
	for _, builtin := range program.Builtins {
		program.RootScope.SymbolTable.Insert(builtin.Symbol)
	}
}

// DefineBuiltinSymbols defines the values in the delegate's scope
func (program *Program) DefineBuiltinValues() {
	// TODO: Define before others
	for _, builtin := range program.Builtins {
		program.Delegate.Define(builtin.Identifier, builtin.Value)
	}
}

func (program *Program) Validate() []error {
	return semantics.Validate(program.Source, program.RootScope)
}

// Run executes a program
func (program *Program) Run() error {
	return program.Engine.Evaluate(program.Source)
}

// RunTask executes a specific task
func (program *Program) RunTask(task string) error {
	// First run the top-level code
	err := program.Engine.Evaluate(program.Source)
	if err != nil {
		return err
	}

	// Run the specified task
	return program.Engine.EvaluateTask(task)
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

func (program *Program) DefineBuiltinValue(identifier string, valueType ValueType, traits semantics.Trait, value interface{}) {
	symbol := semantics.CreateSymbol(identifier, traits, nil)
	program.Builtins[identifier] = &Builtin{
		Identifier: identifier,
		Symbol:     symbol,
		Value:      &Value{Type: valueType, Value: value},
	}
}
