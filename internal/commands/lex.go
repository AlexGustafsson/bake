package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/lexing"
	"github.com/urfave/cli/v2"
)

func lexCommand(context *cli.Context) error {
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
	lexer := lexing.Lex(input)

	for {
		item := lexer.NextItem()
		fmt.Println(item.DebugString())
		if item.Type == lexing.ItemEndOfInput || item.Type == lexing.ItemError {
			break
		}
	}

	return nil
}
