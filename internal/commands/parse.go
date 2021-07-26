package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/parsing"
	"github.com/urfave/cli/v2"
)

func parseCommand(context *cli.Context) error {
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
		return err
	}

	fmt.Print(sourceFile.String())

	return nil
}
