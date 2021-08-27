package main

import (
	"fmt"
	"io/ioutil"
	"os"

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
	program := runtime.CreateProgram(input)
	builtins.Register(program)

	err = program.Parse()
	if err != nil {
		PrintPrettyError(err, input)
		return fmt.Errorf("parsing failed")
	}

	errs := program.BuildSymbols()
	if len(errs) > 0 {
		PrintPrettyErrors(errs, input)
		return fmt.Errorf("validation failed")
	}

	program.DefineBuiltinSymbols()

	errs = program.Validate()
	if len(errs) > 0 {
		PrintPrettyErrors(errs, input)
		return fmt.Errorf("validation failed")
	}

	runtime := runtime.CreateDefaultRuntime()

	program.DefineBuiltinValues(runtime)

	return program.Run(runtime)
}
