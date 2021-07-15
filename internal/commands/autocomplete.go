package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func autocompleteCommand(context *cli.Context) error {
	shell := context.String("shell")

	switch shell {
	case "fish":
		result, err := context.App.ToFishCompletion()
		if err != nil {
			return err
		}

		fmt.Print(result)
	default:
		return fmt.Errorf("unsupported shell '%s'", shell)
	}

	return nil
}
