package parsing

import (
	"github.com/AlexGustafsson/bake/ast"
	"github.com/AlexGustafsson/bake/lexing"
)

func createRangeFromItem(item lexing.Item) *ast.Range {
	start := ast.Position{
		Offset:    item.Range.Start.Offset,
		Line:      item.Range.Start.Line,
		Character: item.Range.Start.Character,
	}

	end := ast.Position{
		Offset:    item.Range.End.Offset,
		Line:      item.Range.End.Line,
		Character: item.Range.End.Character,
	}

	return ast.CreateRange(start, end)
}
