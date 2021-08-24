package state

import (
	"sync"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Document struct {
	sync.Mutex
	URI         string
	Content     string
	Version     int32
	Diagnostics []protocol.Diagnostic
}

func CreateDocument(uri string, content string, version int32) *Document {
	return &Document{
		URI:         uri,
		Content:     content,
		Version:     version,
		Diagnostics: make([]protocol.Diagnostic, 0),
	}
}

func (document *Document) CreateDiagnostic(severity protocol.DiagnosticSeverity, r ast.Range, message string) {
	document.Lock()
	defer document.Unlock()

	source := "bake"

	diagnostic := protocol.Diagnostic{
		Severity: &severity,
		Source:   &source,
		Message:  message,
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      uint32(r.Start().Line),
				Character: uint32(r.Start().Character),
			},
			End: protocol.Position{
				Line:      uint32(r.End().Line),
				Character: uint32(r.End().Character),
			},
		},
	}

	document.Diagnostics = append(document.Diagnostics, diagnostic)
}

func (document *Document) ClearDiagnostics() {
	document.Lock()
	defer document.Unlock()

	document.Diagnostics = make([]protocol.Diagnostic, 0)
}

func (document *Document) PublishDiagnostics(context *glsp.Context) {
	context.Notify("textDocument/publishDiagnostics", protocol.PublishDiagnosticsParams{
		URI:         protocol.DocumentUri(document.URI),
		Diagnostics: document.Diagnostics,
	})
}

func (document *Document) PublishError(context *glsp.Context, err error) {
	context.Notify("window/showMessage", protocol.ShowMessageParams{
		Type:    protocol.MessageTypeError,
		Message: err.Error(),
	})
}
