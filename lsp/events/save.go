package events

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
	"github.com/sourcegraph/go-lsp"
)

func HandleSave(event *Event, document *state.Document) {
	document.ClearDiagnostics()
	defer document.PublishDiagnostics(event.Connection, event.Context)

	source, err := parsing.Parse(document.Content)
	if err != nil {
		if treeError, ok := err.(*ast.TreeError); ok {
			document.CreateDiagnostic(lsp.Error, *treeError.Range, treeError.Message)
		} else {
			document.PublishError(event.Connection, event.Context, err)
		}
		return
	}

	sourceScope, errs := semantics.Build(source)
	if len(errs) > 0 {
		for _, err := range errs {
			if treeError, ok := err.(*ast.TreeError); ok {
				document.CreateDiagnostic(lsp.Error, *treeError.Range, treeError.Message)
			} else {
				document.PublishError(event.Connection, event.Context, err)
			}
		}
		return
	}

	errs = semantics.Validate(source, sourceScope)
	if len(errs) > 0 {
		for _, err := range errs {
			if treeError, ok := err.(*ast.TreeError); ok {
				document.CreateDiagnostic(lsp.Error, *treeError.Range, treeError.Message)
			} else {
				document.PublishError(event.Connection, event.Context, err)
			}
		}
		return
	}
}
