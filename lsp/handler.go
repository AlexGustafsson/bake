package lsp

import (
	"fmt"
	goRuntime "runtime"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/builtins"
	"github.com/AlexGustafsson/bake/lsp/state"
	"github.com/AlexGustafsson/bake/runtime"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Handler struct {
	protocol.Handler
	State *state.State
}

func NewHandler() *Handler {
	handler := &Handler{
		State: state.CreateState(),
	}

	handler.Handler = protocol.Handler{
		Initialize:            handler.handleInitialize,
		Initialized:           handler.handleInitialized,
		TextDocumentDidChange: handler.handleChange,
		TextDocumentDidOpen:   handler.handleOpen,
		TextDocumentDidClose:  handler.handleClose,
		TextDocumentHover:     handler.handleHover,
		TextDocumentDidSave:   handler.handleSave,
		CancelRequest:         handler.handleCancel,
	}

	return handler
}

func (handler *Handler) handleInitialize(context *glsp.Context, parameters *protocol.InitializeParams) (interface{}, error) {
	syncKind := protocol.TextDocumentSyncKindIncremental
	trueOption := true

	return protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: &protocol.TextDocumentSyncOptions{
				OpenClose: &trueOption,
				Save:      &trueOption,
				Change:    &syncKind,
			},
			// CompletionProvider:     &protocol.CompletionOptions{ResolveProvider: &trueOption, TriggerCharacters: []string{"(", "."}},
			// DefinitionProvider:     true,
			// TypeDefinitionProvider: true,
			// DocumentSymbolProvider: true,
			HoverProvider: true,
			// ReferencesProvider:     true,
			// ImplementationProvider: true,
			// SignatureHelpProvider:  &protocol.SignatureHelpOptions{TriggerCharacters: []string{"(", ","}},
		},
	}, nil
}

func (handler *Handler) handleInitialized(context *glsp.Context, parameters *protocol.InitializedParams) error {
	return nil
}

func (handler *Handler) handleOpen(context *glsp.Context, parameters *protocol.DidOpenTextDocumentParams) error {
	if parameters.TextDocument.LanguageID != "bake" {
		return fmt.Errorf("unsupported language")
	}

	handler.State.CreateDocument(string(parameters.TextDocument.URI), parameters.TextDocument.Text, parameters.TextDocument.Version)
	return nil
}

func (handler *Handler) handleClose(context *glsp.Context, parameters *protocol.DidCloseTextDocumentParams) error {
	handler.State.RemoveDocument(string(parameters.TextDocument.URI))
	return nil
}

func (handler *Handler) handleHover(context *glsp.Context, parameters *protocol.HoverParams) (*protocol.Hover, error) {
	return nil, nil
}

func (handler *Handler) handleChange(context *glsp.Context, parameters *protocol.DidChangeTextDocumentParams) (err error) {
	defer handler.recover(&err)

	document := handler.requireDocument(parameters.TextDocument.URI)
	document.Lock()
	defer document.Unlock()

	// Ignore old updates
	if parameters.TextDocument.Version < document.Version {
		return nil
	}

	document.Version = parameters.TextDocument.Version
	for _, change := range parameters.ContentChanges {
		switch cast := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			document.UpdateRange(cast.Range, cast.Text)
		case protocol.TextDocumentContentChangeEventWhole:
			document.Content = cast.Text
		}
	}

	return nil
}

func (handler *Handler) handleSave(context *glsp.Context, parameters *protocol.DidSaveTextDocumentParams) (err error) {
	defer handler.recover(&err)

	document := handler.requireDocument(parameters.TextDocument.URI)

	document.ClearDiagnostics()
	defer document.PublishDiagnostics(context)

	program := runtime.CreateProgram(document.Content, nil)
	builtins.Register(program)

	err = program.Parse()
	if err != nil {
		handler.handleDocumentError(err, document, context)
		return fmt.Errorf("parsing failed")
	}

	errs := program.BuildSymbols()
	if len(errs) > 0 {
		handler.handleDocumentErrors(errs, document, context)
		return fmt.Errorf("validation failed")
	}

	program.DefineBuiltinSymbols()

	errs = program.Validate()
	if len(errs) > 0 {
		handler.handleDocumentErrors(errs, document, context)
		return fmt.Errorf("validation failed")
	}

	return nil
}

func (handler *Handler) handleCancel(context *glsp.Context, parameters *protocol.CancelParams) error {
	return nil
}

func (handler *Handler) handleDocumentError(err error, document *state.Document, context *glsp.Context) {
	if treeError, ok := err.(*ast.TreeError); ok {
		document.CreateDiagnostic(protocol.DiagnosticSeverityError, *treeError.Range, treeError.Message)
	} else {
		document.PublishError(context, err)
	}
}

func (handler *Handler) handleDocumentErrors(errs []error, document *state.Document, context *glsp.Context) {
	for _, err := range errs {
		handler.handleDocumentError(err, document, context)
	}
}

func (handler *Handler) recover(errp *error) {
	err := recover()
	if err != nil {
		if _, ok := err.(goRuntime.Error); ok {
			panic(err)
		}

		if handler != nil {
			*errp = err.(error)
		}
	}
}

// requireDocument handles document retrieval. The document is required to exist
func (handler *Handler) requireDocument(uri string) *state.Document {
	if document, ok := handler.State.Documents[uri]; ok {
		return document
	} else {
		panic(fmt.Errorf("no such document"))
	}
}
