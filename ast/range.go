package ast

import (
	"fmt"
)

// Range describes a range of a text
type Range struct {
	// Start is the inclusive start of the range
	Start Position
	// End is the inclusive end of the range
	End Position
}

// Position describes a position in a text
type Position struct {
	// Line is the zero-based line number
	Line int
	// Character is the zero-based unicode character index
	Character int
	// Offset is the offset from the start of the text, in bytes
	Offset int
}

func (r Range) String() string {
	return fmt.Sprintf("from %s to %s", r.Start, r.End)
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line+1, p.Character+1)
}

func CreateRange(start Position, end Position) *Range {
	return &Range{
		Start: start,
		End:   end,
	}
}
