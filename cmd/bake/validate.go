package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlexGustafsson/bake/builtins"
	"github.com/AlexGustafsson/bake/internal/dot"
	"github.com/AlexGustafsson/bake/runtime"
	"github.com/urfave/cli/v2"
)

func validateCommand(context *cli.Context) error {
	outputType := context.String("type")
	if outputType == "" {
		outputType = "dot"
	}

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

	output := dot.FormatScope(program.MainModule.RootScope)

	if outputType == "dot" {
		fmt.Print(output)
	} else {
		buffer := bytes.NewBufferString(output)
		cmd := exec.Command("dot", "-T", outputType)
		cmd.Stdin = buffer
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
