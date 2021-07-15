package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlexGustafsson/bake/parsing"
	"github.com/AlexGustafsson/bake/parsing/nodes"
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
	tree, err := parsing.Parse(input)
	if err != nil {
		return err
	}

	node := tree.Root
	for node != nil {
		switch node.Type() {
		case nodes.NodeImportType:
			fmt.Println("import")
			fmt.Println("---")
			fmt.Print(node.String())
			fmt.Println("---")
			node = nil
		default:
			fmt.Printf("unsupported type: %d\n", node.Type())
			node = nil
		}
	}
	fmt.Println("end")

	return nil
}
