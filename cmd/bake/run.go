package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/ast"
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

	program.DefineBuiltinFunction("print", -1, func(engine *runtime.Engine, arguments []*runtime.Value) *runtime.Value {
		for i, argument := range arguments {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(argument)
		}
		fmt.Println()
		return nil
	})

	program.DefineBuiltinFunction("range", 2, func(engine *runtime.Engine, arguments []*runtime.Value) *runtime.Value {
		if arguments[0].Type != runtime.ValueTypeNumber || arguments[1].Type != runtime.ValueTypeNumber {
			panic(fmt.Errorf("bad arguments - expected numbers"))
		}

		start := arguments[0].Value.(int)
		stop := arguments[1].Value.(int)

		if stop < start {
			panic(fmt.Errorf("stop cannot be less than start"))
		}

		elements := make([]*runtime.Value, stop-start)
		for i := 0; i < stop-start; i++ {
			elements[i] = &runtime.Value{Type: runtime.ValueTypeNumber, Value: start + i}
		}

		return &runtime.Value{Type: runtime.ValueTypeArray, Value: elements}
	})

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
