package parsing

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/lexing"
)

type ParseError struct {
	Start      int
	Line       int
	Column     int
	TokenValue string
	Message    string
	line       string
}

func CreateParseError(item lexing.Item, line string, format string, args ...interface{}) *ParseError {
	return &ParseError{
		Start:      item.Start,
		Line:       item.Line,
		Column:     item.Column,
		TokenValue: item.Value,
		line:       line,
		Message:    fmt.Sprintf(format, args...),
	}
}

func (parseError *ParseError) Error() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\033[1mfile.bke:%d:%d\033[0m: \033[31;1merror\033[0m: ", parseError.Line, parseError.Column)
	builder.WriteString(parseError.Message)
	builder.WriteRune('\n')

	start := parseError.line[0:parseError.Column]
	end := parseError.line[parseError.Column+len(parseError.TokenValue):]
	fmt.Fprintf(&builder, "%s\033[31;1m%s\033[0m%s\n", start, parseError.TokenValue, end)
	builder.WriteString(strings.Repeat(" ", len(start)))
	fmt.Fprintf(&builder, "\033[31;1m^%s\033[0m\n", strings.Repeat("~", len(parseError.TokenValue)-1))

	return builder.String()
}
