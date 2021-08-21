package events

import (
	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/sourcegraph/go-lsp"
)

func HandleChange(event *Event, document *state.Document) {
	parameters := event.Parameters.(lsp.DidChangeTextDocumentParams)
	document.Lock()
	defer document.Unlock()

	// Ignore old updates
	if parameters.TextDocument.Version < document.Version {
		return
	}

	// TODO: For now, don't support incremental changes
	document.Version = parameters.TextDocument.Version
	if len(parameters.ContentChanges) > 0 {
		document.Content = parameters.ContentChanges[0].Text
	}
}
