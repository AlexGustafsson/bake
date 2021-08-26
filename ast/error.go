package ast

import (
	"fmt"
	"strings"
)

type TreeError struct {
	Message string
	Range   *Range
}

func CreateTreeError(r *Range, format string, args ...interface{}) *TreeError {
	return &TreeError{
		Range:   r,
		Message: fmt.Sprintf(format, args...),
	}
}

func (parseError *TreeError) Error() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\033[1mfile.bke:%d:%d\033[0m: \033[31;1merror\033[0m: ", parseError.Range.Start.Line+1, parseError.Range.Start.Character+1)
	builder.WriteString(parseError.Message)

	return builder.String()
}

func (parseError *TreeError) ErrorWithLine(input string) string {
	var builder strings.Builder

	builder.WriteString(parseError.Error())

	// TODO: make less memory intensive
	allLines := strings.Split(input, "\n")
	lines := make([]string, 0)
	if parseError.Range != nil {
		lines = allLines[parseError.Range.Start.Line : parseError.Range.End.Line+1]
	}

	if parseError.Range.Start.Line == parseError.Range.End.Line {
		builder.WriteRune('\n')
		start := lines[0][0:parseError.Range.Start.Character]
		middle := lines[0][parseError.Range.Start.Character:parseError.Range.End.Character]
		end := lines[0][parseError.Range.End.Character:]
		fmt.Fprintf(&builder, "%s\033[31;1m%s\033[0m%s\n", start, middle, end)
		builder.WriteString(strings.Repeat(" ", len(start)))
		fmt.Fprintf(&builder, "\033[31;1m^%s\033[0m\n", strings.Repeat("~", parseError.Range.End.Character-parseError.Range.Start.Character-1))
	}

	return builder.String()
}
