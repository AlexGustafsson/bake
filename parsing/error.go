package parsing

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/ast"
)

type ParseError struct {
	Message string
	lines   []string
	Range   *ast.Range
}

func CreateParseError(input string, r *ast.Range, format string, args ...interface{}) *ParseError {
	// TODO: make less memory intensive
	allLines := strings.Split(input, "\n")
	lines := make([]string, 0)
	if r != nil {
		lines = allLines[r.Start().Line : r.End().Line+1]
	}

	return &ParseError{
		Range:   r,
		lines:   lines,
		Message: fmt.Sprintf(format, args...),
	}
}

func (parseError *ParseError) Error() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\033[1mfile.bke:%d:%d\033[0m: \033[31;1merror\033[0m: ", parseError.Range.Start().Line+1, parseError.Range.Start().Character+1)
	builder.WriteString(parseError.Message)
	builder.WriteRune('\n')

	if parseError.Range.Start().Line == parseError.Range.End().Line {
		start := parseError.lines[0][0:parseError.Range.Start().Character]
		middle := parseError.lines[0][parseError.Range.Start().Character:parseError.Range.End().Character]
		end := parseError.lines[0][parseError.Range.End().Character:]
		fmt.Fprintf(&builder, "%s\033[31;1m%s\033[0m%s\n", start, middle, end)
		builder.WriteString(strings.Repeat(" ", len(start)))
		fmt.Fprintf(&builder, "\033[31;1m^%s\033[0m\n", strings.Repeat("~", parseError.Range.End().Character-parseError.Range.Start().Character-1))
	}

	return builder.String()
}
