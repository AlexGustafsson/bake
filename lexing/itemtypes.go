package lexing

type ItemType int

const (
	ItemSpace ItemType = iota
	ItemNewline
	ItemEndOfFile
	ItemError
)
