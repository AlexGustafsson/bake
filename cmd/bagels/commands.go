package main

import (
	"github.com/urfave/cli/v2"
)

// commands contains all commands of the application
var commands = []*cli.Command{
	{
		Name:   "version",
		Usage:  "Show the application's version",
		Action: versionCommand,
	},
	{
		Name:   "start",
		Usage:  "Start the server",
		Action: startCommand,
	},
}
