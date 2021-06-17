package commands

import (
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
	parser := parsing.Parse(input)
	parser.Wait()

	return nil
}
