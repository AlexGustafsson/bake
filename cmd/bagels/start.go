package main

import (
	"github.com/AlexGustafsson/bake/lsp"
	"github.com/tliron/glsp/server"
	cli "github.com/urfave/cli/v2"

	// Must include a backend implementation. See kutil's logging/ for other options.
	"github.com/tliron/kutil/logging"
	_ "github.com/tliron/kutil/logging/simple"
)

func startCommand(_ *cli.Context) error {
	logging.Configure(1, nil)

	handler := lsp.NewHandler()
	server := server.NewServer(handler, "bake", false)
	server.RunStdio()
	return nil
}
