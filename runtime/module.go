package runtime

import (
	"strings"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
)

// Module is a single unit of a Bake source file
type Module struct {
	Input     string
	Source    *ast.Block
	RootScope *semantics.Scope
	// Package is the module's package, if any
	Package string
	imports map[string]bool
}

func CreateModule(input string) *Module {
	return &Module{
		Input:   input,
		imports: make(map[string]bool),
	}
}

// Parse parses the source of the module
func (module *Module) Parse() error {
	source, err := parsing.Parse(module.Input)
	if err != nil {
		return err
	}

	module.Source = source

	module.populateImports()
	module.populatePackage()

	return nil
}

func (module *Module) populateImports() {
	for _, node := range module.Source.Statements {
		if importsDeclaration, ok := node.(*ast.ImportsDeclaration); ok {
			for _, path := range importsDeclaration.Imports {
				// Assume only string parts (valid after semantic validation has been performed)
				var builder strings.Builder
				for _, part := range path.Parts {
					if stringPart, ok := part.(*ast.StringPart); ok {
						builder.WriteString(stringPart.Content)
					}
				}
				resolvedPath := builder.String()
				module.imports[resolvedPath] = true
			}
		}
	}
}

func (module *Module) populatePackage() {
	for _, node := range module.Source.Statements {
		if packageDeclaration, ok := node.(*ast.PackageDeclaration); ok {
			module.Package = packageDeclaration.Identifier
		}
	}
}

// BuildSymbols builds the root semantic scope of the parsed module
func (module *Module) BuildSymbols() []error {
	rootScope, errs := semantics.Build(module.Source)
	if len(errs) > 0 {
		return errs
	}

	module.RootScope = rootScope
	return []error{}
}

// Validate validates the semantics of the module
func (module *Module) Validate() []error {
	return semantics.Validate(module.Source, module.RootScope)
}

// Evaluates evaluates a module
func (module *Module) Evaluate(engine *Engine, delegate Delegate) error {
	return engine.Evaluate(module.Source)
}

// Imports are all of the imported paths
func (module *Module) Imports() []string {
	paths := make([]string, 0)
	for path := range module.imports {
		paths = append(paths, path)
	}
	return paths
}
