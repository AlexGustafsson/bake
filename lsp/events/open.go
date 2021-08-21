package events

import (
	"fmt"

	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/sourcegraph/go-lsp"
)

func HandleOpen(event *Event, state *state.State) error {
	parameters := event.Parameters.(lsp.DidOpenTextDocumentParams)

	if parameters.TextDocument.LanguageID != "bake" {
		return fmt.Errorf("unsupported language")
	}

	state.CreateDocument(string(parameters.TextDocument.URI), parameters.TextDocument.Text, parameters.TextDocument.Version)
	return nil
}
