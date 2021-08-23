package events

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/sourcegraph/go-lsp"
)

func HandleSave(event *Event, document *state.Document) {
	document.ClearDiagnostics()

	_, err := parsing.Parse(document.Content)
	if err != nil {
		if treeError, ok := err.(*ast.TreeError); ok {
			document.CreateDiagnostic(lsp.Error, *treeError.Range, treeError.Message)
		} else {
			document.PublishError(event.Connection, event.Context, err)
		}
	}

	document.PublishDiagnostics(event.Connection, event.Context)
}
