package lsp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AlexGustafsson/bake/parsing"
	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler Handler) Handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) {
	response, err := handler.handle(ctx, connection, request)
	if err != nil {
		log.Error("error creating response", err)
		return
	}

	err = connection.Reply(ctx, request.ID, response)
	if err != nil {
		log.Error("error sending response", err)
	}
}

// https://github.com/a-h/qt-lsp/blob/6c6909dc8457bdd14ee3c80513c0cc87394c6aa7/handler.go#L18
func (handler Handler) handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (interface{}, error) {
	switch request.Method {
	case "initialize":
		if request.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}

		var params lsp.InitializeParams
		if err := json.Unmarshal(*request.Params, &params); err != nil {
			return nil, err
		}

		log.Info("starting in '%s'", params.RootURI)

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
		}, nil
	case "initialized":
		// A notification that the client is ready to receive requests. Ignore
		return nil, nil
	case "exit":
		connection.Close()
		return nil, nil
	case "shutdown":
		return nil, nil
	case "textDocument/hover":
		if request.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}
		var params lsp.TextDocumentPositionParams
		if err := json.Unmarshal(*request.Params, &params); err != nil {
			return nil, err
		}

		return nil, nil
	case "textDocument/didOpen":
		if request.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}

		var params lsp.DidOpenTextDocumentParams
		if err := json.Unmarshal(*request.Params, &params); err != nil {
			return nil, err
		}

		if params.TextDocument.LanguageID != "bake" {
			return nil, fmt.Errorf("unsupported language")
		}

		_, err := parsing.Parse(params.TextDocument.Text)
		if err != nil {
			if parseError, ok := err.(*parsing.ParseError); ok {
				connection.Notify(ctx, "textDocument/publishDiagnostics", lsp.PublishDiagnosticsParams{
					URI: params.TextDocument.URI,
					Diagnostics: []lsp.Diagnostic{
						{
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      parseError.Line,
									Character: parseError.Column,
								},
								End: lsp.Position{
									Line:      parseError.Line,
									Character: parseError.Column + len(parseError.TokenValue),
								},
							},
							Severity: lsp.Error,
							Source:   "bake",
							Message:  parseError.Message,
						},
					},
				})
			} else {
				connection.Notify(ctx, "window/showMessage", lsp.ShowMessageParams{
					Type:    lsp.MTError,
					Message: err.Error(),
				})
			}
		}

		return nil, nil
	case "textDocument/didClose":
		if request.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}

		var params lsp.DidOpenTextDocumentParams
		if err := json.Unmarshal(*request.Params, &params); err != nil {
			return nil, err
		}

		return nil, nil
	case "textDocument/didChange":
		if request.Params == nil {
			return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
		}

		var params lsp.DidChangeTextDocumentParams
		if err := json.Unmarshal(*request.Params, &params); err != nil {
			return nil, err
		}

		for _, change := range params.ContentChanges {
			log.Info("from %d:%d to %d:%d: %s", change.Range.Start.Line, change.Range.Start.Character, change.Range.End.Line, change.Range.End.Character, change.Text)
		}

		return nil, nil
	}

	return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", request.Method)}
}
