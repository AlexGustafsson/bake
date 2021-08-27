package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/builtins"
	"github.com/AlexGustafsson/bake/runtime"
	"github.com/urfave/cli/v2"
)

func runCommand(context *cli.Context) error {
	inputPath := context.String("input")
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	input := string(inputBytes)
	program := runtime.CreateProgram(input, runtime.CreateDefaultRuntime())
	builtins.Register(program)

	errs := program.Run()
	if len(errs) > 0 {
		for _, err := range errs {
			if treeError, ok := err.(*ast.TreeError); ok {
				// Print the formatted error
				fmt.Fprint(os.Stderr, treeError.ErrorWithLine(input))
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		return fmt.Errorf("invalid program")
	}

	return nil
}
