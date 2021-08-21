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
		Name:   "parse",
		Usage:  "Parse a Bakefile",
		Action: parseCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Path to input file",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Usage:       "Output type",
				DefaultText: "dot",
			},
		},
	},
	{
		Name:   "format",
		Usage:  "Format a Bakefile",
		Action: formatCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Path to input file",
				Required: true,
			},
		},
	},
	{
		Name:   "lex",
		Usage:  "Lex a Bakefile",
		Action: lexCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Path to input file",
				Required: true,
			},
		},
	},
	{
		Name:    "init",
		Aliases: []string{"initialize"},
		Usage:   "Initialize a bake",
		Action:  initCommand,
	},
	{
		Name:   "autocomplete",
		Usage:  "Print autocompletion for a shell",
		Action: autocompleteCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "shell",
				Usage:    "Name of the shell to print an autocompletion script for",
				Required: true,
			},
		},
	},
}
