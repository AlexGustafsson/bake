package main

import (
	"fmt"
	"io/ioutil"
	"os"

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
	program, errs := runtime.CreateProgram(input)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return fmt.Errorf("invalid program")
	}

	engine := runtime.CreateEngine(runtime.CreateDefaultRuntime())
	err = engine.Evaluate(program)
	if err != nil {
		return err
	}

	return nil
}
