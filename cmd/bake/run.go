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
	if inputPath == "" {
		inputPath = "Bakefile"
	}

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

	program.DefineBuiltinValues()

	err = program.Run()
	if err != nil {
		return err
	}

	for _, task := range context.Args().Slice() {
		err = program.RunTask(task)
		if err != nil {
			return err
		}
	}

	return nil
}
