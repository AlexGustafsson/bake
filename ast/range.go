package ast

import (
	"fmt"
)

// Range describes a range of a text
type Range struct {
	// start is the inclusive start of the range
	start Position
	// snd is the inclusive end of the range
	end Position
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
	return fmt.Sprintf("from %s to %s", r.start, r.end)
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line+1, p.Character+1)
}

func (r Range) Start() Position {
	return r.start
}

func (r Range) End() Position {
	return r.end
}

func CreateRange(start Position, end Position) Range {
	return Range{
		start: start,
		end:   end,
	}
}
