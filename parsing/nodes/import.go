package nodes

import (
	"strings"
)

type NodeImport struct {
	NodeType
	NodePosition
	Imports []string
}

func (node *NodeImport) String() string {
	var builder strings.Builder
	builder.WriteString("import (\n")
	for _, imported := range node.Imports {
		builder.WriteString("  ")
		builder.WriteString(imported)
		builder.WriteByte('\n')
	}
	builder.WriteString(")\n")
	return builder.String()
}
