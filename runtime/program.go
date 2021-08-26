package runtime

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
)

type Program struct {
	Source    *ast.SourceFile
	RootScope *semantics.Scope
}

func CreateProgram(input string) (*Program, []error) {
	source, err := parsing.Parse(input)
	if err != nil {
		return nil, []error{err}
	}

	rootScope, errs := semantics.Build(source)
	if len(errs) > 0 {
		return nil, errs
	}

	errs = semantics.Validate(source, rootScope)
	if len(errs) > 0 {
		return nil, errs
	}

	return &Program{
		Source:    source,
		RootScope: rootScope,
	}, nil
}
