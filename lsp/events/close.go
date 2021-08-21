package events

import (
	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/sourcegraph/go-lsp"
)

func HandleClose(event *Event, state *state.State) {
	parameters := event.Parameters.(lsp.DidCloseTextDocumentParams)
	state.RemoveDocument(string(parameters.TextDocument.URI))
}
