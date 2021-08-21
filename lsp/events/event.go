package events

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
)

type Event struct {
	Context    context.Context
	Connection *jsonrpc2.Conn
	Parameters interface{}
}

func CreateEvent(ctx context.Context, connection *jsonrpc2.Conn, parameters interface{}) *Event {
	return &Event{
		Context:    ctx,
		Connection: connection,
		Parameters: parameters,
	}
}
