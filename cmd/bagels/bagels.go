package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexGustafsson/bake/lsp"
	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/jsonrpc2"
)

type stdstream struct{}

func (stdstream) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (stdstream) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (stdstream) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}

	return os.Stdout.Close()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	handler := lsp.NewHandler()

	stream := jsonrpc2.NewBufferedStream(stdstream{}, jsonrpc2.VSCodeObjectCodec{})
	rpcLogger := jsonrpc2.LogMessages(log.New())
	connection := jsonrpc2.NewConn(ctx, stream, handler, rpcLogger)
	select {
	case <-ctx.Done():
		log.Info("signal received")
		connection.Close()
	case <-connection.DisconnectNotify():
		log.Info("client disconnected")
	}

	log.Info("stopped")
}
