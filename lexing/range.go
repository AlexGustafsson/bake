package lexing

import "fmt"

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
	return fmt.Sprintf("from %d:%d to %d:%d", r.Start.Line+1, r.Start.Character+1, r.End.Line+1, r.End.Character+1)
}
