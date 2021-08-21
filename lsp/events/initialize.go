package events

import (
	"github.com/sourcegraph/go-lsp"
)

func HandleInitialize(event *Event) lsp.InitializeResult {
	// TODO: Switch to incremental when there is support
	// Send full text on each update
	kind := lsp.TDSKFull

	return lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
				Kind: &kind,
			},
			CompletionProvider:     &lsp.CompletionOptions{ResolveProvider: true, TriggerCharacters: []string{"(", "."}},
			DefinitionProvider:     true,
			TypeDefinitionProvider: true,
			DocumentSymbolProvider: true,
			HoverProvider:          true,
			ReferencesProvider:     true,
			ImplementationProvider: true,
			SignatureHelpProvider:  &lsp.SignatureHelpOptions{TriggerCharacters: []string{"(", ","}},
		},
	}
}
