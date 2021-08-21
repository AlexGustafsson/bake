package state

import (
	"context"
	"sync"

	"github.com/AlexGustafsson/bake/parsing/nodes"
	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

type Document struct {
	sync.Mutex
	URI         string
	Content     string
	Version     int
	Diagnostics []lsp.Diagnostic
}

func CreateDocument(uri string, content string, version int) *Document {
	return &Document{
		URI:         uri,
		Content:     content,
		Version:     version,
		Diagnostics: make([]lsp.Diagnostic, 0),
	}
}

func (document *Document) CreateDiagnostic(severity lsp.DiagnosticSeverity, r nodes.Range, message string) {
	document.Lock()
	defer document.Unlock()

	diagnostic := lsp.Diagnostic{
		Severity: severity,
		Source:   "bake",
		Message:  message,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      r.Start().Line,
				Character: r.Start().Character,
			},
			End: lsp.Position{
				Line:      r.End().Line,
				Character: r.End().Character,
			},
		},
	}

	document.Diagnostics = append(document.Diagnostics, diagnostic)
}

func (document *Document) ClearDiagnostics() {
	document.Lock()
	defer document.Unlock()

	document.Diagnostics = make([]lsp.Diagnostic, 0)
}

func (document *Document) PublishDiagnostics(connection *jsonrpc2.Conn, ctx context.Context) {
	connection.Notify(ctx, "textDocument/publishDiagnostics", lsp.PublishDiagnosticsParams{
		URI:         lsp.DocumentURI(document.URI),
		Diagnostics: document.Diagnostics,
	})
}

func (document *Document) PublishError(connection *jsonrpc2.Conn, ctx context.Context, err error) {
	connection.Notify(ctx, "window/showMessage", lsp.ShowMessageParams{
		Type:    lsp.MTError,
		Message: err.Error(),
	})
}
