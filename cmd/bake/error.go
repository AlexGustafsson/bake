package main

import (
	"fmt"
	"os"

	"github.com/AlexGustafsson/bake/ast"
)

func PrintPrettyError(err error, input string) {
	if treeError, ok := err.(*ast.TreeError); ok {
		fmt.Fprintf(os.Stderr, treeError.ErrorWithLine(input))
	} else {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}

func PrintPrettyErrors(errs []error, input string) {
	for _, err := range errs {
		PrintPrettyError(err, input)
	}
}
