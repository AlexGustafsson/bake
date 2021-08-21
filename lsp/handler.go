package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/AlexGustafsson/bake/lsp/events"
	"github.com/AlexGustafsson/bake/lsp/state"
	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
)

type Handler struct {
	State *state.State
}

func NewHandler() *Handler {
	return &Handler{
		State: state.CreateState(),
	}
}

// Handle is the main entrypoint for events
func (handler Handler) Handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) {
	// Call the internal handle function and send any potential error or reply
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

// handle is the internal handler for determining the correct action for an event
func (handler Handler) handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (reply interface{}, err error) {
	defer handler.recover(&err)

	switch request.Method {
	case "initialize":
		var params lsp.InitializeParams
		handler.requireParameters(request, &params)

		event := events.CreateEvent(ctx, connection, params)
		response := events.HandleInitialize(event)
		return response, nil
	case "initialized":
		// A notification that the client is ready to receive requests. Ignore
		return nil, nil
	case "exit":
		connection.Close()
		return nil, nil
	case "shutdown":
		return nil, nil
	case "textDocument/hover":
		var params lsp.TextDocumentPositionParams
		handler.requireParameters(request, &params)

		event := events.CreateEvent(ctx, connection, params)
		events.HandleHover(event)
		return nil, nil
	case "textDocument/didOpen":
		var params lsp.DidOpenTextDocumentParams
		handler.requireParameters(request, &params)

		event := events.CreateEvent(ctx, connection, params)
		err = events.HandleOpen(event, handler.State)
		return nil, err
	case "textDocument/didClose":
		var parameters lsp.DidCloseTextDocumentParams
		handler.requireParameters(request, &parameters)

		event := events.CreateEvent(ctx, connection, parameters)
		events.HandleClose(event, handler.State)
		return nil, nil
	case "textDocument/didChange":
		var params lsp.DidChangeTextDocumentParams
		handler.requireParameters(request, &params)

		document := handler.requireDocument(string(params.TextDocument.URI))

		event := events.CreateEvent(ctx, connection, params)
		events.HandleChange(event, document)
		return nil, nil
	case "textDocument/didSave":
		var params lsp.DidSaveTextDocumentParams
		handler.requireParameters(request, &params)

		document := handler.requireDocument(string(params.TextDocument.URI))

		event := events.CreateEvent(ctx, connection, params)
		events.HandleSave(event, document)
		return nil, nil
	}

	return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", request.Method)}
}

func (handler *Handler) recover(errp *error) {
	err := recover()
	if err != nil {
		if _, ok := err.(runtime.Error); ok {
			panic(err)
		}

		if handler != nil {
			*errp = err.(error)
		}
	}
}

// requireParameters handles unmarshalling of a request's parameter that must not be nil.
// the parameters argument is expected to be a pointer to a value of the type to unmarshal
func (handler *Handler) requireParameters(request *jsonrpc2.Request, parameters interface{}) {
	if request.Params == nil {
		panic(&jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams})
	}

	if err := json.Unmarshal(*request.Params, parameters); err != nil {
		panic(err)
	}
}

// requireDocument handles document retrieval. The document is required to exist
func (handler *Handler) requireDocument(uri string) *state.Document {
	if document, ok := handler.State.Documents[uri]; ok {
		return document
	} else {
		panic(&jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams, Message: "no such document"})
	}
}
