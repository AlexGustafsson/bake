package dot

import (
	"fmt"
	"strings"

	"github.com/AlexGustafsson/bake/semantics"
)

func FormatScope(scope *semantics.Scope) string {
	var builder strings.Builder

	builder.WriteString("digraph G {\n")
	builder.WriteString("splines = false;\n")
	// Global scope
	builder.WriteString(formatScope(scope, "global scope"))
	// Child scopes
	for _, child := range scope.ChildScopes {
		builder.WriteString(formatScope(child, "top-level scope"))
	}
	builder.WriteString("}\n")

	return builder.String()
}

func formatScope(scope *semantics.Scope, name string) string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "subgraph \"cluster_%p\" {\n", scope)
	fmt.Fprintf(&builder, "label=\"%s\";\n", name)
	fmt.Fprintf(&builder, "\"%p\" [shape=none, label=<%s>];\n", scope, formatSymbolTable(scope.SymbolTable))
	builder.WriteString("}\n")

	for _, child := range scope.ChildScopes {
		fmt.Fprintf(&builder, "\"%p\" -> \"%p\";\n", scope, child)
		builder.WriteString(formatScope(child, "child scope"))
	}

	return builder.String()
}

func formatSymbolTable(table *semantics.SymbolTable) string {
	var builder strings.Builder
	builder.WriteString(`<table border="0" cellborder="1" cellspacing="0" cellpadding="4">`)
	builder.WriteString("\n<tr><td>Name</td><td>Trait</td><td>Line</td><td>Argument Count</td></tr>\n")
	for _, symbol := range table.Symbols() {
		traits := strings.Join(symbol.Trait.Strings(), ",")
		fmt.Fprintf(&builder, "<tr><td>%s</td><td>%s</td><td>%d</td><td>%d</td></tr>\n", symbol.Name, traits, symbol.Node.Start().Line+1, symbol.ArgumentCount)
	}
	builder.WriteString("</table>")
	return builder.String()
}
