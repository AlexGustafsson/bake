package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlexGustafsson/bake/parsing"
	"github.com/urfave/cli/v2"
)

func parseCommand(context *cli.Context) error {
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
		return err
	}

	output := sourceFile.DotString()

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
