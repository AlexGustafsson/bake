package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/internal/dot"
	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/semantics"
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
	sourceFile, err := parsing.Parse(input)
	if err != nil {
		if treeError, ok := err.(*ast.TreeError); ok {
			// Print the formatted error
			fmt.Fprint(os.Stderr, treeError.ErrorWithLine(input))
			return fmt.Errorf("parsing failed")
		} else {
			return err
		}
	}

	rootScope, errs := semantics.Build(sourceFile)
	if len(errs) > 0 {
		// Print the formatted errors
		for _, err := range errs {
			if treeError, ok := err.(*ast.TreeError); ok {
				// Print the formatted error
				fmt.Fprint(os.Stderr, treeError.ErrorWithLine(input))
			} else {
				fmt.Fprint(os.Stderr, err)
			}
		}
		return fmt.Errorf("validation failed")
	}

	errs = semantics.Validate(sourceFile, rootScope)
	if len(errs) > 0 {
		// Print the formatted errors
		for _, err := range errs {
			if treeError, ok := err.(*ast.TreeError); ok {
				// Print the formatted error
				fmt.Fprint(os.Stderr, treeError.ErrorWithLine(input))
			} else {
				fmt.Fprint(os.Stderr, err)
			}
		}
		return fmt.Errorf("validation failed")
	}

	output := dot.FormatScope(rootScope)

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
