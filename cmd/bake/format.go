package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/parsing"
	"github.com/urfave/cli/v2"
)

func formatCommand(context *cli.Context) error {
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
		// Print the formatted error
		fmt.Fprint(os.Stderr, err)
		return fmt.Errorf("parsing failed")
	}

	fmt.Print(sourceFile.String())

	return nil
}
